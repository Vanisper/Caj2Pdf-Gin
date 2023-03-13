#!/usr/bin/bash

function copy_folder {
    local source_folder=$1
    local target_folder=$2

    mkdir -p "$target_folder/$source_folder"
    #echo "$target_folder/$source_folder"
    cp -rf "$source_folder/." "$target_folder/$source_folder"
}

function is_windows() {
    if [[ $(uname) == "MINGW"* || $(uname) == "CYGWIN"* ]]; then
        return 0 # 返回 true
    else
        return 1 # 返回 false
    fi
}

echo 前端安装依赖 ...
cd ./templates && pnpm install && cd ..

echo 前端项目构建 ...
cd ./templates && pnpm run build && cd ..

echo Gin项目分发 ...
if is_windows; then
    goreleaser release --snapshot --clean && chown -R "$(whoami):$(id -g)" ./dist
else
    sudo goreleaser release --snapshot --clean && sudo chown -R "$(whoami):$(id -gn)" ./dist
fi

echo 复制依赖文件 ...
source_folders=("templates/dist" "lib")
# set lib 777
if is_windows; then
    chmod -R 777 "lib"
else
    sudo chmod -R 777 "lib"
fi


prefix="Caj2Pdf-Gin_"

# shellcheck disable=SC2207
target_folders=($(find ./dist -maxdepth 1 -type d -name "${prefix}*"))

for target_folder in "${target_folders[@]}"; do
    for source_folder in "${source_folders[@]}"; do
        copy_folder "$source_folder" "$target_folder"
    done
done


