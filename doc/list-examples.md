# List Examples<a name="EN-US_TOPIC_0142723183"></a>

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
```

Based on the structure of objects in the bucket, different object listing scenarios require different commands.

-   To list three objects in bucket  **bucket-test**, the command is as follows:

    ```
    ./obsutil ls obs://bucket-test -limit=3
    ```

    The returned result is listed in lexicographical order by object name and version ID as follows:

    ```
    obs://bucket-test/test1.txt
    obs://bucket-test/test2.txt
    obs://bucket-test/test3.txt
    ```

-   To list three objects following  **test3.txt**  in bucket  **bucket-test**, the command is as follows:

    ```
    ./obsutil ls obs://bucket-test -limit=3 -marker=test3.txt
    ```

    The returned result is listed in lexicographical order by object name and version ID as follows:

    ```
    obs://bucket-test/test4.txt
    obs://bucket-test/test5.txt
    obs://bucket-test/test6.txt
    ```

-   To list the files and subdirectories in the root directory of bucket  **bucket-test**  in non-recursive mode, that is, files in the subdirectories are not listed, the command is as follows:

    ```
    ./obsutil ls obs://bucket-test -d
    ```

    The returned result is listed in lexicographical order by object name and version ID as follows:

    ```
    obs://bucket-test/test1.txt
    obs://bucket-test/test2.txt
    obs://bucket-test/test3.txt
    obs://bucket-test/test4.txt
    obs://bucket-test/test5.txt
    obs://bucket-test/test6.txt
    obs://bucket-test/src1/
    obs://bucket-test/src2/
    ```


>![](public_sys-resources/icon-note.gif) **NOTE:**   
>All the commands in the preceding examples use the Linux OS as the running environment.  

