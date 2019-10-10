# Creating a Folder<a name="EN-US_TOPIC_0160351845"></a>

## Function<a name="section1479112110815"></a>

You can use this command to create a folder in a specified bucket or local file system.

## Command Line Structure<a name="section1220752192216"></a>

-   In Windows
    -   Creating a folder in a specified bucket

        ```
        obsutil mkdir obs://bucket/folder[/subfolder1/subfolder2] [-config=xxx]
        ```

    -   Creating a folder in the local file system

        ```
        obsutil mkdir folder_url [-config=xxx]
        ```


-   In Linux or macOS
    -   Creating a folder in a specified bucket

        ```
        ./obsutil mkdir obs://bucket/folder[/subfolder1/subfolder2] [-config=xxx]
        ```

    -   Creating a folder in the local file system

        ```
        ./obsutil mkdir folder_url [-config=xxx]
        ```



## Parameter Description<a name="section6559191102418"></a>

<a name="table10831182114445"></a>
<table><thead align="left"><tr id="row683212154419"><th class="cellrowborder" valign="top" width="13%" id="mcps1.1.4.1.1"><p id="p118329219446"><a name="p118329219446"></a><a name="p118329219446"></a>Parameter</p>
</th>
<th class="cellrowborder" valign="top" width="18%" id="mcps1.1.4.1.2"><p id="p15137125919108"><a name="p15137125919108"></a><a name="p15137125919108"></a>Optional or Mandatory</p>
</th>
<th class="cellrowborder" valign="top" width="69%" id="mcps1.1.4.1.3"><p id="p12832121184414"><a name="p12832121184414"></a><a name="p12832121184414"></a>Description</p>
</th>
</tr>
</thead>
<tbody><tr id="row1096715662419"><td class="cellrowborder" valign="top" width="13%" headers="mcps1.1.4.1.1 "><p id="p11357141182414"><a name="p11357141182414"></a><a name="p11357141182414"></a>bucket</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.2 "><p id="p12357511172416"><a name="p12357511172416"></a><a name="p12357511172416"></a>Mandatory when creating a folder in a specified bucket</p>
</td>
<td class="cellrowborder" valign="top" width="69%" headers="mcps1.1.4.1.3 "><p id="p1435717116242"><a name="p1435717116242"></a><a name="p1435717116242"></a>Bucket name</p>
</td>
</tr>
<tr id="row129424415242"><td class="cellrowborder" valign="top" width="13%" headers="mcps1.1.4.1.1 "><p id="p03571611152415"><a name="p03571611152415"></a><a name="p03571611152415"></a>folder</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.2 "><p id="p15150103020252"><a name="p15150103020252"></a><a name="p15150103020252"></a>Mandatory when creating a folder in a specified bucket</p>
</td>
<td class="cellrowborder" valign="top" width="69%" headers="mcps1.1.4.1.3 "><p id="p5378637132511"><a name="p5378637132511"></a><a name="p5378637132511"></a>Folder path in the bucket. This value can contain multi-level folders. Separate each level with a slash (/).</p>
</td>
</tr>
<tr id="row108328217449"><td class="cellrowborder" valign="top" width="13%" headers="mcps1.1.4.1.1 "><p id="p64495172515"><a name="p64495172515"></a><a name="p64495172515"></a>folder_url</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.2 "><p id="p154316502519"><a name="p154316502519"></a><a name="p154316502519"></a>Mandatory when creating a folder in the local file system</p>
</td>
<td class="cellrowborder" valign="top" width="69%" headers="mcps1.1.4.1.3 "><p id="p17425512259"><a name="p17425512259"></a><a name="p17425512259"></a>Folder path in the local file system. The value can be an absolute path or a relative path.</p>
</td>
</tr>
<tr id="row1492014154359"><td class="cellrowborder" valign="top" width="13%" headers="mcps1.1.4.1.1 "><p id="p153951131317"><a name="p153951131317"></a><a name="p153951131317"></a>config</p>
</td>
<td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.2 "><p id="p12395135316"><a name="p12395135316"></a><a name="p12395135316"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="69%" headers="mcps1.1.4.1.3 "><p id="p43952034313"><a name="p43952034313"></a><a name="p43952034313"></a>User-defined configuration file for executing a command. For details about parameters that can be configured, see <a href="parameter-description.md">Parameter Description</a>.</p>
</td>
</tr>
</tbody>
</table>

## Running Example<a name="section15899161919244"></a>

-   Take the Windows OS as an example. Run the  **obsutil mkdir obs://bucket-test/folder1/folder2**  command to create a folder in a bucket.

```
obsutil mkdir obs://bucket-test/folder1/folder2

Create folder [obs://bucket-test/folder1/] successfully, request id [0000016979E1D23C860BB3D8E4577C5E]
Create folder [obs://bucket-test/folder1/folder2] successfully, request id [0000016979E1D2B2860BB5181229C72C]
```

