# Synchronous Upload Examples<a name="EN-US_TOPIC_0150889364"></a>

Assume that a local folder is in the following structure:

```
└── src1
    ├── src2
        ├── test1.txt
        └── test2.txt
    ├── src3
    └── test3.txt
```

Assume that bucket  **bucket-test**  contains the following objects:

```
obs://bucket-test/src1/
obs://bucket-test/src1/src2/
obs://bucket-test/src1/src2/test1.txt
obs://bucket-test/src1/src3/
```

Based on the structure of the preceding local folder and objects in the bucket, different synchronous upload scenarios require different commands.

-   To synchronize the  **test3.txt**  file in the local  **src1**  folder to the root directory of bucket  **bucket-test**, the command is as follows:

    ```
    ./obsutil sync /src1/test3.txt  obs://bucket-test
    ```

    After the synchronization is successful, the  **test3.txt**  file is directly uploaded to the bucket because there is no  **test3.txt**  in bucket  **bucket-test**. Then, objects in the bucket are as follows:

    ```
    obs://bucket-test/test3.txt
    obs://bucket-test/src1/
    obs://bucket-test/src1/src2/
    obs://bucket-test/src1/src2/test1.txt
    obs://bucket-test/src1/src3/
    ```

-   To synchronize all files and subfolders in the local  **src1**  folder to the  **src**  folder in bucket  **bucket-test**, the command is as follows:

    ```
    ./obsutil sync /src1  obs://bucket-test/src1
    ```

    After the synchronization, the objects in the bucket are as follows:

    ```
    obs://bucket-test/src1/
    obs://bucket-test/src1/test3.txt
    obs://bucket-test/src1/src2/
    obs://bucket-test/src1/src2/test1.txt
    obs://bucket-test/src1/src2/test2.txt
    obs://bucket-test/src1/src3/
    ```


>![](public_sys-resources/icon-note.gif) **NOTE:**   
>All the commands in the preceding examples use the Linux OS as the running environment.  

