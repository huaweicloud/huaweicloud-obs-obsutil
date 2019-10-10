# Synchronous Download Examples<a name="EN-US_TOPIC_0150889365"></a>

Assume that bucket  **bucket-test**  contains the following objects:

```
obs://bucket-test/src1/
obs://bucket-test/src1/test3.txt
obs://bucket-test/src1/src2/
obs://bucket-test/src1/src2/test1.txt
obs://bucket-test/src1/src2/test2.txt
obs://bucket-test/src1/src3/
```

Assume that a local folder is in the following structure:

```
└── src1
    └── test3.txt
```

Based on the structure of the preceding local folder and objects in the bucket, different synchronous download scenarios require different commands.

-   To synchronize all files and subfolders in the  **src1**  folder in bucket  **bucket-test**  to the local  **src1**  folder, the command is as follows:

    ```
    ./obsutil sync obs://bucket-test/src1  /src1
    ```

    After the synchronization is successful, the following files are generated in the local  **src1**  folder:

    ```
    └── src1
        ├── src2
            ├── test1.txt
            └── test2.txt
        ├── src3
        └── test3.txt
    ```


>![](public_sys-resources/icon-note.gif) **NOTE:**   
>All the commands in the preceding examples use the Windows OS as the running environment.  

