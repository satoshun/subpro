#!/bin/sh

root=/usr/local/bin

cp ./subpro ${root}/
chmod 755 ${root}/subpro

cp ./subpro_completion ${root}/__subpro_completion
chmod 755 ${root}/__subpro_completion
