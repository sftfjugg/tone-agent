#!/bin/bash

# set -x

DIR_TMP="/tmp/toneagent_install"

toneagnet_deb_x86="toneagent-1.0.3-x86_64.deb"
toneagnet_deb_arm="toneagent-1.0.3-aarch64.deb"
toneagnet_rpm_x86="toneagent-1.0.3-1.an8.x86_64.rpm"
toneagnet_rpm_arm="toneagent-1.0.3-1.an8.aarch64.rpm"

toneagent_file="toneagent.tar.gz"

if [ -d $DIR_TMP ]; then
    # echo $DIR_TMP "exists"
    rm -rf $DIR_TMP
fi

mkdir $DIR_TMP

ARCHIVE=$(awk '/^__ARCHIVE_BOUNDARY__/ { print NR + 1; exit 0; }' $0)

tail -n +$ARCHIVE $0 > $DIR_TMP/$toneagent_file
tar -zpxf $DIR_TMP/$toneagent_file -C $DIR_TMP/

check_sys(){
    local checkType=$1
    local value=$2

    local release=''
    local systemPackage=''
    local arch=''

    if [[ -f /etc/redhat-release ]]; then
        release='centos'
        systemPackage='yum'
    elif grep -Eqi 'debian|raspbian' /etc/issue; then
        release='debian'
        systemPackage='apt'
    elif grep -Eqi 'ubuntu' /etc/issue; then
        release='ubuntu'
        systemPackage='apt'
    elif grep -Eqi 'centos|red hat|redhat' /etc/issue; then
        release='centos'
        systemPackage='yum'
    elif grep -Eqi 'debian|raspbian' /proc/version; then
        release='debian'
        systemPackage='apt'
    elif grep -Eqi 'ubuntu' /proc/version; then
        release='ubuntu'
        systemPackage='apt'
    elif grep -Eqi 'centos|red hat|redhat' /proc/version; then
        release='centos'
        systemPackage='yum'
    fi

    if [[ "$systemPackage" == "" ]];then
        if type apt > /dev/null 2>&1; then
            systemPackage='apt'
        elif type yum > /dev/null 2>&1; then
            systemPackage='yum'
        fi
    fi

    case $(uname -m) in
        x86_64)  arch='x86';;
        aarch64) arch='arm';;
    esac

    if [[ "${checkType}" == 'sysRelease' ]]; then
        if [ "${value}" == "${release}" ]; then
            return 0
        else
            return 1
        fi
    elif [[ "${checkType}" == 'packageManager' ]]; then
        if [ "${value}" == "${systemPackage}" ]; then
            return 0
        else
            return 1
        fi
    elif [[ "${checkType}" == 'architecture' ]]; then
        if [ "${value}" == "${arch}" ]; then
            return 0
        else
            return 1
        fi
    fi
}

if check_sys packageManager yum; then
    if check_sys architecture x86; then
        rpm -i $DIR_TMP/file/$toneagnet_rpm_x86
    elif check_sys architecture arm; then
        rpm -i $DIR_TMP/file/$toneagnet_rpm_arm
    fi

elif check_sys packageManager apt; then
    if check_sys architecture x86; then
        dpkg -i $DIR_TMP/file/$toneagnet_deb_x86
    elif check_sys architecture arm; then
        dpkg -i $DIR_TMP/file/$toneagnet_deb_arm
    fi
fi

exit 0
__ARCHIVE_BOUNDARY__
