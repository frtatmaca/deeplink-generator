#!/bin/bash  
#!/bin/bash


RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # no color
CHECK='✅'
CROSS='❌'
function app(){
 	cd /d/workspace/go/src/gitlab.zfu.fintek.local/zfu/collateral-app/
}
function api(){
	cd /d/workspace/go/src/gitlab.zfu.fintek.local/zfu/collaterals-api/
}
function bff(){
	cd /d/workspace/go/src/gitlab.zfu.fintek.local/zfu/bffcollateral/
}

function fed(){
	cd /d/workspace/fed/$1/$2/$3/$4
}

function gvtupdate(){
	grep \"gitlab.*$1 vendor/manifest | awk '{ gsub(",", "", $2); print $2 }' | xargs --no-run-if-empty -L 1 gvt update -precaire 
}

function fedCommit() {
	CURRENT_BRANCH=$(git branch --show-current)
	if [[ "${CURRENT_BRANCH}" == "dev" ]]; then
		read -p "Enter Commit Message: " MESSAGE
		git add --all && git commit -m "${MESSAGE}" "${@}" && git pull --no-edit && git push && git checkout master && git merge dev --no-edit && git pull --no-edit && git push && git checkout dev
	else
		echo "You are not on branch dev, current branch is ${CURRENT_BRANCH}"
	fi	
}

function fedCommitSelf(){
	read -p "Enter Commit Message: " MESSAGE
	git add --all && git commit -m "${MESSAGE}" "${@}" && git pull --no-edit && git push
}

# deletes node modules
function cleanNodeModules(){
	echo -e -n "${YELLOW}Cleaning up node_modules...${NC}"
	temp=$(rm -rf node_modules)
	if [[ $? -eq 0 ]]
	then
		echo -e "${GREEN}${CHECK}${NC}";
	else
		echo -e "${RED}${CROSS}Due to ${temp}${NC}"
	fi;
	return 0;
}
# deletes node modules
function cleanYarnLock(){
	echo -e -n "${YELLOW}Cleaning up yarn.lock...${NC}"
	rm yarn.lock 2> /dev/null
	echo -e "${GREEN}${CHECK}${NC}"	
}

# deletes node modules
function cleanYarnCache(){
	echo -e -n "${YELLOW}Cleaning up yarn cache...${NC}"
	temp=$(yarn cache clean 1> /dev/null)
	if [[ $? -eq 0 ]]
	then
		echo -e "${GREEN}${CHECK}${NC}";
	else
		echo -e "${RED}${CROSS}Failed due to ${temp}${NC} "
	fi;
	return 0;
}
# deletes node modules
function installDeps(){
	echo -e "${YELLOW}Installing depedencies...${NC}"
	yarn install
	if [[ $? -eq 0 ]]
	then
		echo -e "${GREEN}DONE...${CHECK}${NC}";
	else
		echo -e "${RED}Failed...${CROSS}${NC} "
	fi;
	return 0;
}

function fresh(){
	echo "Do you want to clean up yarn cache?";
	
	select yn in "Yes" "No"; do
		case ${yn} in
		Yes ) 
			cleanNodeModules &&
			cleanYarnLock &&
			cleanYarnCache &&
			installDeps &&
			return 0;
		;;
		No ) 
			cleanNodeModules &&
			cleanYarnLock &&
			installDeps &&
			return 0;
		;;
		esac;
	done;
}

function personalFns (){
	echo "app: goes to collateral-app directory";
	echo "api: goes to collateral-api directory";
	echo "bff: goes to bffcollateral directory";
	echo "fed: goes to fed base directory, usage `fed exp exptalep`, to go project directory of given module followed by project name";
	echo "gvtupdate: updates dependency files";
	echo "fedCommit: runnable on branch dev of fed project, commits, merges origin dev, then merges local dev to master, merges origin master then return dev branch again";
	echo "fedCommitSelf: commit, pull, push of current branch.";
	echo "cleanBuild: clean build current fed project";
}

export -f app
export -f api
export -f bff
export -f fed
export -f gvtupdate
export -f fedCommit
export -f fedCommitSelf
export -f fresh
export -f personalFns
 
 
