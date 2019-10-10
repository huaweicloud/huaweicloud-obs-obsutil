# Download Examples<a name="EN-US_TOPIC_0142723181"></a>

Assume that bucket  **bucket-test**  contains the following objects:

```
obs://bucket-test/test1.txt
obs://bucket-test/test2.txt
obs://bucket-test/test3.txt
obs://bucket-test/test4.txt
obs://bucket-test/test5.txt
obs://bucket-test/test6.txt
obs://bucket-test/src1/
obs://bucket-test/src1/test7.txt
obs://bucket-test/src2/
obs://bucket-test/src2/test8.txt
obs://bucket-test/src2/src3/
obs://bucket-test/src2/src3/test9.txt
```

Based on the structure of objects in the bucket, different download scenarios require different commands.

-   To download the  **test1.txt**  file from bucket  **bucket-test**  to the local  **src1**  folder, the command is as follows:

    ```
    ./obsutil cp obs://bucket-test/test1.txt  /src1
    ```

    After the download is complete, the following file is generated on the local PC:

    ```
    └── src1
        └── test1.txt
    ```

-   Run the following command to download the  **test1.txt**  file to your local PC. If there is no  **test.txt**  on the local PC, the  **test1.txt**  file is directly downloaded and you can rename it to  **test.txt**. If  **test.txt**  already exists,  **test1.txt**  is downloaded and overwrites the original local  **test.txt**  file after renaming.

    ```
    ./obsutil cp obs://bucket-test/test1.txt  /test.txt
    ```

    After the download is complete, the following file generated on the local PC:

    ```
    └── test.txt
    ```

-   To recursively download the entire  **src2**  folder from bucket  **bucket-test**  to the local  **src1**  folder in force mode, the command is as follows:

    ```
    ./obsutil cp obs://bucket-test/src2  /src1 -r -f
    ```

    After the download is complete, the following files are generated on the local PC:

    ```
    └── src1
        └── src2
            ├── src3
                └── test9.txt
            └── test8.txt
    ```

-   To recursively download all files and subfolders in the  **src2**  folder from bucket  **bucket-test**  to the local  **src1**  folder in force mode, the command is as follows:

    ```
    ./obsutil cp obs://bucket-test/src2  /src1 -r -f -flat
    ```

    After the download is complete, the following files are generated on the local PC:

    ```
    └── src1
        ├── src3
            └── test9.txt
        └── test8.txt
    ```

-   To recursively download the all objects in bucket  **bucket-test**  to the local  **src0**  folder in force mode, the command is as follows:

    ```
    ./obsutil cp obs://bucket-test  /src0 -r -f
    ```

    After the download is complete, the following files are generated on the local PC:

    ```
    └── src0
        ├── test1.txt
        ├── test2.txt
        ├── test3.txt
        ├── test4.txt
        ├── test5.txt
        ├── test6.txt
        ├── src1
            └── test7.txt
        └── src2
            ├── src3
                └── test9.txt
            └── test8.txt
    ```


>![](public_sys-resources/icon-note.gif) **NOTE:**   
>All the commands in the preceding examples use the Linux OS as the running environment.  

