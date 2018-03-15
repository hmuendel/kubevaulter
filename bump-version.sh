#!/usr/bin/env bash -euo pipefail

#default values
version_file=./version
version_upgrade=patch
minor=false
patch=true
write=false

#usage explanation
usage="$(basename "$0")  [-f filepath] [-h] [-v version] [-w]  -- this script returns the increased semversion stored in a file

where:
    -f path to version file containing a single semversion default: $version_file
    -h show this help text
    -v version to bump default: $version_upgrade
    -w  overwrite the the original version file with the bumped version"

# reading flags
while getopts ':f:hv:w' option; do
  case "$option" in
    f) version_file="$OPTARG"
       ;;
    h) echo "$usage"
       exit 0
       ;;
    v) version_upgrade="$OPTARG"
       ;;
    w) write=true
       ;;
    :) printf "missing argument for -%s\n" "$OPTARG" >&2
       echo "$usage" >&2
       exit 1
       ;;
   \?) printf "illegal option: -%s\n" "$OPTARG" >&2
       echo "$usage" >&2
       exit 1
       ;;
  esac
done


#reading version from file
version=$(cat $version_file)

#parsing seperate parts of semversion
major=$(echo $version | awk -F'.' '{print $1}')
minor=$(echo $version | awk -F'.' '{print $2}')
patch=$(echo $version | awk -F'.' '{print $3}')


#increasing version according to flags
case "$version_upgrade" in
    major) ((major++))
           minor=0
           patch=0
           ;;
    minor) ((minor++))
           patch=0
           ;;
    patch) ((patch++))
           ;;
esac

# write new version to file, only if write flag is set
if [ "$write" = "true" ]; then
    echo $major.$minor.$patch > version
fi
# output bumped version
echo $major.$minor.$patch