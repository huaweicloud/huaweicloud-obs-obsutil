# Examples of Listing Multipart Upload Tasks<a name="EN-US_TOPIC_0142725103"></a>

Assume that bucket  **bucket-test**  contains the following multipart upload tasks:

```
obs://bucket-test/task1.txt uploadid1
obs://bucket-test/task1.txt uploadid2
obs://bucket-test/task2.txt uploadid3
obs://bucket-test/task3.txt uploadid4
obs://bucket-test/src1/
obs://bucket-test/src1/task4.txt uploadid5
obs://bucket-test/src2/
obs://bucket-test/src2/task5.txt uploadid6
```

-   Run the following command to list three multipart upload tasks in bucket  **bucket-test**:

    ```
    ./obsutil ls obs://bucket-test -m -limit=3
    ```

    The returned result is listed in lexicographical order by object name as follows:

    ```
    obs://bucket-test/task1.txt uploadid1
    obs://bucket-test/task1.txt uploadid2
    obs://bucket-test/task2.txt uploadid3
    ```

-   To list the rest multipart upload tasks following  **uploadid1**, the command is as follows:

    ```
    obsutil ls obs://bucket-test -m -limit=3 -marker=task1.txt -uploadIdMarker=uploadid1
    ```

    The returned result is listed in lexicographical order by object name and upload ID as follows:

    ```
    obs://bucket-test/task1.txt uploadid2
    obs://bucket-test/task2.txt uploadid3
    obs://bucket-test/task3.txt uploadid4
    ```


>![](public_sys-resources/icon-note.gif) **NOTE:**   
>All the commands in the preceding examples use the Linux OS as the running environment.  

