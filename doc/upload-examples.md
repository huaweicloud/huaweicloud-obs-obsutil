# Upload Examples<a name="EN-US_TOPIC_0142723120"></a>

Assume that a local folder is in the following structure:

```
└── src1
    ├── src2
        ├── test1.txt
        └── test2.txt
    ├── src3
    └── test3.txt
```

Based on the preceding folder structure, different upload scenarios require different commands.

-   To upload the  **test3.txt**  file in the local  **src1**  folder to the root directory of bucket  **bucket-test**, the command is as follows:

    ```
    ./obsutil cp /src1/test3.txt  obs://bucket-test
    ```

    After the upload is successful, the following object is generated in the bucket:

    ```
    ./obs://bucket-test/test3.txt
    ```

-   To upload the  **test3.txt**  file in the local  **src1**  folder to the root directory of bucket  **bucket-test**  and rename it to  **aaa.txt**, the command is as follows:

    ```
    ./obsutil cp /src1/test3.txt  obs://bucket-test/aaa.txt
    ```

    After the upload is successful, the following object is generated in the bucket:

    ```
    ./obs://bucket-test/aaa.txt
    ```

-   To upload the  **test3.txt**  file in the local  **src1**  folder to the  **src**  folder in bucket  **bucket-test**, the command is as follows:

    ```
    ./obsutil cp /src1/test3.txt  obs://bucket-test/src/
    ```

    After the upload is successful, the following object is generated in the bucket:

    ```
    ./obs://bucket-test/src/test3.txt
    ```

-   To recursively upload the entire local  **src2**  folder to the root directory of bucket  **bucket-test**  in force mode, the command is as follows:

    ```
    ./obsutil cp /src1/src2  obs://bucket-test -r -f
    ```

    After the upload is successful, the following objects are generated in the bucket:

    ```
    obs://bucket-test/src2/
    obs://bucket-test/src2/test1.txt
    obs://bucket-test/src2/test2.txt
    ```

-   To recursively upload the entire local  **src1**  folder to the  **src**  folder in bucket  **bucket-test**  in force mode, the command is as follows:

    ```
    ./obsutil cp /src1  obs://bucket-test/src -r -f
    ```

    After the upload is successful, the following objects are generated in the bucket:

    ```
    obs://bucket-test/src/src1/
    obs://bucket-test/src/src1/src2/
    obs://bucket-test/src/src1/src2/test1.txt
    obs://bucket-test/src/src1/src2/test2.txt
    obs://bucket-test/src/src1/src3/
    obs://bucket-test/src/src1/test3.txt
    ```

-   To recursively upload the all files and subfolders in the local  **src1**  folder to the  **src**  folder in bucket  **bucket-test**  in force mode, the command is as follows:

    ```
    ./obsutil cp /src1  obs://bucket-test/src -r -f -flat
    ```

    After the upload is successful, the following objects are generated in the bucket:

    ```
    obs://bucket-test/src/
    obs://bucket-test/src/src2/
    obs://bucket-test/src/src2/test1.txt
    obs://bucket-test/src/src2/test2.txt
    obs://bucket-test/src/src3/
    obs://bucket-test/src/test3.txt
    ```


>![](public_sys-resources/icon-note.gif) **NOTE:**   
>All the commands in the preceding examples use the Linux OS as the running environment.  

