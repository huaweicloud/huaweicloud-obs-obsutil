# Deleting a Bucket<a name="EN-US_TOPIC_0141181365"></a>

## Function<a name="section1479112110815"></a>

You can use this command to delete a bucket. The bucket to be deleted must be empty \(containing no objects, historical versions, or fragments\).

>![](public_sys-resources/icon-note.gif) **NOTE:**   
>To delete a non-empty bucket, run the commands in  [Deleting a Multipart Upload Task](deleting-a-multipart-upload-task.md)  and in  [Deleting an Object](deleting-an-object.md)  to clear the bucket, and then run the following command to delete the bucket.  

## Command Line Structure<a name="section1220752192216"></a>

-   In Windows

    ```
    obsutil rm obs://bucket [-f] [-config=xxx]
    ```

-   In Linux or macOS

    ```
    ./obsutil rm obs://bucket [-f] [-config=xxx]
    ```


## Parameter Description<a name="section6559191102418"></a>

<a name="table10831182114445"></a>
<table><thead align="left"><tr id="row683212154419"><th class="cellrowborder" valign="top" width="16.16161616161616%" id="mcps1.1.4.1.1"><p id="p118329219446"><a name="p118329219446"></a><a name="p118329219446"></a>Parameter</p>
</th>
<th class="cellrowborder" valign="top" width="28.28282828282828%" id="mcps1.1.4.1.2"><p id="p15137125919108"><a name="p15137125919108"></a><a name="p15137125919108"></a>Optional or Mandatory</p>
</th>
<th class="cellrowborder" valign="top" width="55.55555555555556%" id="mcps1.1.4.1.3"><p id="p12832121184414"><a name="p12832121184414"></a><a name="p12832121184414"></a>Description</p>
</th>
</tr>
</thead>
<tbody><tr id="row108328217449"><td class="cellrowborder" valign="top" width="16.16161616161616%" headers="mcps1.1.4.1.1 "><p id="p64495172515"><a name="p64495172515"></a><a name="p64495172515"></a>bucket</p>
</td>
<td class="cellrowborder" valign="top" width="28.28282828282828%" headers="mcps1.1.4.1.2 "><p id="p154316502519"><a name="p154316502519"></a><a name="p154316502519"></a>Mandatory</p>
</td>
<td class="cellrowborder" valign="top" width="55.55555555555556%" headers="mcps1.1.4.1.3 "><p id="p17425512259"><a name="p17425512259"></a><a name="p17425512259"></a>Bucket name</p>
</td>
</tr>
<tr id="row107241436114819"><td class="cellrowborder" valign="top" width="16.16161616161616%" headers="mcps1.1.4.1.1 "><p id="p5724123684813"><a name="p5724123684813"></a><a name="p5724123684813"></a>f</p>
</td>
<td class="cellrowborder" valign="top" width="28.28282828282828%" headers="mcps1.1.4.1.2 "><p id="p1272453612483"><a name="p1272453612483"></a><a name="p1272453612483"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.55555555555556%" headers="mcps1.1.4.1.3 "><p id="p972417362487"><a name="p972417362487"></a><a name="p972417362487"></a>Runs in force mode.</p>
</td>
</tr>
<tr id="row1660212336344"><td class="cellrowborder" valign="top" width="16.16161616161616%" headers="mcps1.1.4.1.1 "><p id="p153951131317"><a name="p153951131317"></a><a name="p153951131317"></a>config</p>
</td>
<td class="cellrowborder" valign="top" width="28.28282828282828%" headers="mcps1.1.4.1.2 "><p id="p12395135316"><a name="p12395135316"></a><a name="p12395135316"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.55555555555556%" headers="mcps1.1.4.1.3 "><p id="p43952034313"><a name="p43952034313"></a><a name="p43952034313"></a>User-defined configuration file for executing a command. For details about parameters that can be configured, see <a href="parameter-description.md">Parameter Description</a>.</p>
</td>
</tr>
</tbody>
</table>

## Running Example<a name="section15899161919244"></a>

-   Take the Windows OS as an example. Run the  **obsutil rm obs://bucket-test**  command to delete bucket  **bucket-test**.

```
obsutil rm obs://bucket-test

Do you want delete bucket [bucket-test] ? Please input (y/n) to confirm:
y
Delete bucket [bucket-test] successfully!
```

