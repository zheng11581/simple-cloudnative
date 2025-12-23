FROM swr.cn-north-4.myhuaweicloud.com/pixiu-public/centos:centos8.1.1911

COPY ./mem-alloc/mem_alloc /
COPY ./read-file/read_file /

CMD ["/read_file", "/mnt/test.file", "100"]