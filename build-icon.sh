#!/bin/bash

# KeyView 图标生成脚本
# 需要安装 ImageMagick: brew install imagemagick

APP_NAME="KeyView"
SVG_FILE="icon-template.svg"
ICONSET_DIR="${APP_NAME}.iconset"
ICON_FILE="build/darwin/icons.icns"

echo "🎨 生成 KeyView 图标..."

# 创建 iconset 目录
rm -rf ${ICONSET_DIR}
mkdir -p ${ICONSET_DIR}

# 生成所需的所有尺寸（优化压缩）
echo "📏 生成不同尺寸..."
sizes=(16 32 128 256 512 1024)
for size in "${sizes[@]}"; do
    echo "  生成 ${size}x${size}..."
    convert -background none -density 72 -colorspace sRGB -strip ${SVG_FILE}[0] -resize ${size}x${size} PNG32:${ICONSET_DIR}/icon_${size}x${size}.png

    if [ $size -lt 1024 ]; then
        double_size=$((size * 2))
        echo "  生成 ${double_size}x${double_size}..."
        convert -background none -density 72 -colorspace sRGB -strip ${SVG_FILE}[0] -resize ${double_size}x${double_size} PNG32:${ICONSET_DIR}/icon_${size}x${size}@2x.png
    fi
done

# 生成 icns 文件
echo "📦 生成 .icns 文件..."
iconutil -c icns ${ICONSET_DIR} -o ${ICON_FILE}

# 显示文件大小
FILE_SIZE=$(ls -lh ${ICON_FILE} | awk '{print $5}')
echo "📊 文件大小: ${FILE_SIZE}"

# 清理临时文件
rm -rf ${ICONSET_DIR}

echo "✅ 图标生成完成: ${ICON_FILE}"
