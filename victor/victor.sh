#!/bin/ksh
################################################################################
#
# Victor, the cleaner.
#
# Deletes files, e.g. old log files. The default match pattern is *.log.
# Cleaning happens in one of two modes, either based on the timestamp of the
# file or based an a total size for a given match pattern.
#
# Date based sample invocation:
#
#  ./victor.sh -f -p "*_appX_*.log" -d 30
#
# will delete files matching the pattern that are older than 30 days.
#
# Size based sample invocation.
#
#  ./victor.sh -f -p "*_appX_*.log" -s 100
#
# will repeatedly delete the oldest file matching the pattern until the total
# disk size is below the specified amount in MB, in this sample 100 MB.
#
# If, for any reasons, you want to bypass deletion of certain files or
# directories (e.g. if they are under investigation right now, and you want to
# make sure they survive and wipes), just 'touch' them.
#
# Limits:
#
# - This script will remove a maximum of 1023 files due an internal ksh
#   limitation.
# - size based deletion does not work on nested/ deep directories
#
################################################################################

################################################################################
# Function section
################################################################################

usage() {
  echo >&2 \
    "usage: $0 [-f] [-v] [-p pattern] (-d days|-s size in MB) directory"
    exit 1
}

################################################################################
# Main section
################################################################################

[ ! -z "${DEBUG}" ] && set -x

days=0
size=0
force=
verbose=off
pattern="*.log"

while getopts d:fvp:s: opt
do
  case "${opt}" in
      d) days="${OPTARG}";;
      f) force=on;;
      v) verbose=on;;
      p) pattern="${OPTARG}";;
      s) size="${OPTARG}";;
      \?) # Unknown flag
          usage;;
  esac
done

if [ ! -z "${verbose}" ]; then
    echo "Using pattern ${pattern}"
fi

shift `expr ${OPTIND} - 1`
directory="$1"

if [ -z "${directory}" ]; then
  usage
fi

# Check for required parameter days or size
if [ 0 -eq "${days}" -a 0 -eq "${size}" ]; then
    usage
fi

# Ready to rumble...
if [ "${days}" -gt 0 ]; then
    if [ ! -z "${verbose}" ]; then
        echo "Entering time mode"
    fi
    limited=
    index=0
    find "${directory}" -name "${pattern}" -mtime +"${days}" | {
        while read line; do
            # echo "Marking ${line} for deletion"
            files[index]="${line}"
            index=`expr ${index} + 1`

            # According to the spec, ksh88 array limit is 1023
            if [ 1023 -eq ${index} ]; then
                limited=on
                break
            fi
        done
    }

    # Do not go for 'rm ${files[*]}' because this might break size of commandline expansion
    for file in ${files[*]}; do
        if [ ! -z "${verbose}" ]; then
            echo "Deleting $file"
        fi
        if [ ! -z "${force}" ]; then
            rm -rf "${file}"
        fi
    done
    if [ ! -z "${verbose}" ]; then
        echo "Deleted ${index} files"
    fi
    if [ ! -z "${limited}" ]; then
        echo "During execution, we've hit an internal limit. Re-running this command will free more resources."
    fi
elif [ "${size}" -gt 0 ]; then
    if [ ! -z "${verbose}" ]; then
        echo "Entering size mode"
    fi

    # ls -s reports block size, and du does not allow patterns :-(
    currentsize=$(ls -l "${directory}/"${pattern}|awk 'BEGIN { sum=0 } { sum = sum + $5 } END { print sum }')
    mb=`expr ${currentsize} / 1024 / 1024`
    del=`expr ${mb} - ${size}`
    echo "Initial size: ${mb} MB"

    while [ ${del} -gt 0 ]; do
        echo "Free'ing another ${del} MB"
        index=0
        # Delete oldest logfile
        ls -tr "${directory}/"${pattern} | {
            read line
            if [ -f "${line}" ]; then
                echo "Deleting ${line}"
                if [ ! -z "${force}" ]; then
                    rm -f "${line}"
                fi
            else
                # No more files for pattern found
                echo "No more files matching pattern ${pattern} found. Nothing else to delete."
                break
            fi
        }
        currentsize=$(ls -l "${directory}/"${pattern}|awk 'BEGIN { sum=0 } { sum = sum + $5 } END { print sum }')
        mb=`expr ${currentsize} / 1024 / 1024`
        del=`expr ${mb} - ${size}`
        echo "Current size: ${mb} MB"
    done
fi

if [ -z "${force}" ]; then
    echo "Running in dryrun mode. Nothing has been changed. Use -f to really delete."
fi

################################################################################
# EOF
################################################################################
