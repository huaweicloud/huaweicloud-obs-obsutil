# Copy Examples<a name="EN-US_TOPIC_0142723182"></a>

Assume that bucket  **bucket-src**  contains the following objects:

```
obs://bucket-src/test1.txt
obs://bucket-src/test2.txt
obs://bucket-src/test3.txt
obs://bucket-src/test4.txt
obs://bucket-src/test5.txt
obs://bucket-src/test6.txt
obs://bucket-src/src1/
obs://bucket-src/src1/test7.txt
obs://bucket-src/src2/
obs://bucket-src/src2/test8.txt
obs://bucket-src/src2/src3/
obs://bucket-src/src2/src3/test9.txt
```

Based on the structure of objects in the bucket, different copy scenarios require different commands.

-   To copy the  **test1.txt**  file from bucket  **bucket-src**  to bucket  **bucket-dest**, the command is as follows:

    ```
    ./obsutil cp obs://bucket-src/test1.txt  obs://bucket-dest
    ```

    After the copy is complete, the following object is generated in bucket  **bucket-dest**:

    ```
    obs://bucket-dest/test1.txt
    ```

-   To copy the content of the  **test1.txt**  file in bucket  **bucket-src**  to the  **text.txt**  file in bucket  **bucket-dest**, the command is as follows:

    ```
    ./obsutil cp obs://bucket-src/test1.txt  obs://bucket-dest/test.txt
    ```

    After the copy is complete, the following object is generated in bucket  **bucket-dest**:

    ```
    obs://bucket-dest/test.txt
    ```

-   To copy the  **test1.txt**  file in bucket  **bucket-src**  to the  **text**  folder in bucket  **bucket-dest**, the command is as follows:

    ```
    ./obsutil cp obs://bucket-src/test1.txt  obs://bucket-dest/test/
    ```

    After the copy is complete, the following object is generated in bucket  **bucket-dest**:

    ```
    obs://bucket-dest/test/test1.txt
    ```

-   Run the following command to recursively copy the entire  **src2**  folder in bucket  **bucket-src**  to bucket  **bucket-dest**  in force mode:

    ```
    ./obsutil cp obs://bucket-src/src2  obs://bucket-dest -r -f
    ```

    After the copy is complete, the following objects are generated in bucket  **bucket-dest**:

    ```
    obs://bucket-dest/src2/
    obs://bucket-dest/src2/test8.txt
    obs://bucket-dest/src2/src3/
    obs://bucket-dest/src2/src3/test9.txt
    ```

-   To recursively copy all files and subfolders in the  **src2**  folder in bucket  **bucket-src**  to bucket  **bucket-dest**  in force mode, the command is as follows:

    ```
    ./obsutil cp obs://bucket-src/src2  obs://bucket-dest -r -f -flat
    ```

    After the copy is complete, the following objects are generated in bucket  **bucket-dest**:

    ```
    obs://bucket-dest/test8.txt
    obs://bucket-dest/src3/
    obs://bucket-dest/src3/test9.txt
    ```


>![](public_sys-resources/icon-note.gif) **NOTE:**   
>All the commands in the preceding examples use the Linux OS as the running environment.  

