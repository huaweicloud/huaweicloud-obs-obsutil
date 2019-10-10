# Synchronous Copy Examples<a name="EN-US_TOPIC_0150889366"></a>

Assume that the source bucket  **bucket-src**  contains the following objects:

```
obs://bucket-src/src1/
obs://bucket-src/src1/test3.txt
obs://bucket-src/src1/src2/
obs://bucket-src/src1/src2/test1.txt
obs://bucket-src/src1/src2/test2.txt
obs://bucket-src/src1/src3/
```

Assume that the destination bucket  **bucket-dest**  contains the following objects:

```
obs://bucket-dest/src1/
obs://bucket-dest/src1/test3.txt
```

Based on the structure of objects in the bucket, different synchronous copy scenarios require different commands.

-   To synchronize all files and subfolders in the  **src1**  folder in bucket  **bucket-src**  to the  **src1**  folder in bucket  **bucket-dest**, the command is as follows:

    ```
    ./obsutil sync obs://bucket-src/src1  obs://bucket-dest/src1
    ```

    After the synchronous copy is complete, the objects in the destination bucket  **bucket-dest**  are as follows:

    ```
    obs://bucket-dest/src1/
    obs://bucket-dest/src1/test3.txt
    obs://bucket-dest/src1/src2/
    obs://bucket-dest/src1/src2/test1.txt
    obs://bucket-dest/src1/src2/test2.txt
    obs://bucket-dest/src1/src3/
    ```


>![](public_sys-resources/icon-note.gif) **NOTE:**   
>All the commands in the preceding examples use the Windows OS as the running environment.  

