# Listing Multipart Upload Tasks<a name="EN-US_TOPIC_0142356367"></a>

## Function<a name="section12810504216"></a>

You can use this command to query multipart upload tasks in a bucket.

## Command Line Structure<a name="section315294314215"></a>

-   In Windows

    ```
    obsutil ls obs://bucket[/prefix] [-s] [-d] -m [-a] [-uploadIdMarker=xxx] [-marker=xxx] [-limit=1] [-config=xxx]
    ```

-   In Linux or macOS

    ```
    ./obsutil ls obs://bucket[/prefix] [-s] [-d] -m [-a] [-uploadIdMarker=xxx] [-marker=xxx] [-limit=1] [-config=xxx]
    ```


## Parameter Description<a name="section358316614229"></a>

<a name="table11357111919227"></a>
<table><thead align="left"><tr id="row19450419112217"><th class="cellrowborder" valign="top" width="18.181818181818183%" id="mcps1.1.4.1.1"><p id="p9450619132218"><a name="p9450619132218"></a><a name="p9450619132218"></a>Parameter</p>
</th>
<th class="cellrowborder" valign="top" width="25.252525252525253%" id="mcps1.1.4.1.2"><p id="p1045051914221"><a name="p1045051914221"></a><a name="p1045051914221"></a>Optional or Mandatory</p>
</th>
<th class="cellrowborder" valign="top" width="56.565656565656575%" id="mcps1.1.4.1.3"><p id="p34505193226"><a name="p34505193226"></a><a name="p34505193226"></a>Description</p>
</th>
</tr>
</thead>
<tbody><tr id="row1145014197226"><td class="cellrowborder" valign="top" width="18.181818181818183%" headers="mcps1.1.4.1.1 "><p id="p4450141912217"><a name="p4450141912217"></a><a name="p4450141912217"></a>bucket</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p19450219142210"><a name="p19450219142210"></a><a name="p19450219142210"></a>Mandatory</p>
</td>
<td class="cellrowborder" valign="top" width="56.565656565656575%" headers="mcps1.1.4.1.3 "><p id="p18452619122215"><a name="p18452619122215"></a><a name="p18452619122215"></a>Bucket name</p>
</td>
</tr>
<tr id="row745261918223"><td class="cellrowborder" valign="top" width="18.181818181818183%" headers="mcps1.1.4.1.1 "><p id="p345261932219"><a name="p345261932219"></a><a name="p345261932219"></a>prefix</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p345211982219"><a name="p345211982219"></a><a name="p345211982219"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="56.565656565656575%" headers="mcps1.1.4.1.3 "><p id="p1445231992218"><a name="p1445231992218"></a><a name="p1445231992218"></a>Object name prefix for listing multipart uploads</p>
<div class="note" id="note172548512206"><a name="note172548512206"></a><a name="note172548512206"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p7254145112017"><a name="p7254145112017"></a><a name="p7254145112017"></a>If this parameter is left blank, all multipart upload tasks in the bucket are listed.</p>
</div></div>
</td>
</tr>
<tr id="row1845231962214"><td class="cellrowborder" valign="top" width="18.181818181818183%" headers="mcps1.1.4.1.1 "><p id="p14452919202212"><a name="p14452919202212"></a><a name="p14452919202212"></a>s</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p9452161922213"><a name="p9452161922213"></a><a name="p9452161922213"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="56.565656565656575%" headers="mcps1.1.4.1.3 "><p id="p12452161914224"><a name="p12452161914224"></a><a name="p12452161914224"></a>Displays simplified query result.</p>
<div class="note" id="note1452085114212"><a name="note1452085114212"></a><a name="note1452085114212"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p252035122111"><a name="p252035122111"></a><a name="p252035122111"></a>In the simplified format, the returned result contains only the object name and upload ID of the multipart upload.</p>
</div></div>
</td>
</tr>
<tr id="row14452219102219"><td class="cellrowborder" valign="top" width="18.181818181818183%" headers="mcps1.1.4.1.1 "><p id="p1345291912217"><a name="p1345291912217"></a><a name="p1345291912217"></a>d</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p245215195224"><a name="p245215195224"></a><a name="p245215195224"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="56.565656565656575%" headers="mcps1.1.4.1.3 "><p id="p1445218197221"><a name="p1445218197221"></a><a name="p1445218197221"></a>Lists only the multipart upload tasks and sub-directories in the current directory are listed, instead of recursively listing all the multipart upload tasks and sub-directories.</p>
</td>
</tr>
<tr id="row154291518173013"><td class="cellrowborder" valign="top" width="18.181818181818183%" headers="mcps1.1.4.1.1 "><p id="p1645215197222"><a name="p1645215197222"></a><a name="p1645215197222"></a>m</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p18452019162217"><a name="p18452019162217"></a><a name="p18452019162217"></a>Mandatory (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="56.565656565656575%" headers="mcps1.1.4.1.3 "><p id="p745214195227"><a name="p745214195227"></a><a name="p745214195227"></a>Lists multipart upload tasks in the bucket.</p>
</td>
</tr>
<tr id="row87391618304"><td class="cellrowborder" valign="top" width="18.181818181818183%" headers="mcps1.1.4.1.1 "><p id="p19452121932218"><a name="p19452121932218"></a><a name="p19452121932218"></a>a</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p0452141914224"><a name="p0452141914224"></a><a name="p0452141914224"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="56.565656565656575%" headers="mcps1.1.4.1.3 "><p id="p114521319172214"><a name="p114521319172214"></a><a name="p114521319172214"></a>Lists the objects and the multipart upload tasks in the bucket.</p>
</td>
</tr>
<tr id="row127451045183012"><td class="cellrowborder" valign="top" width="18.181818181818183%" headers="mcps1.1.4.1.1 "><p id="p17557916141117"><a name="p17557916141117"></a><a name="p17557916141117"></a>marker</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p1155818161117"><a name="p1155818161117"></a><a name="p1155818161117"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="56.565656565656575%" headers="mcps1.1.4.1.3 "><p id="p5558816151116"><a name="p5558816151116"></a><a name="p5558816151116"></a>Indicates the upload ID after which the multipart upload listing begins. All returned multipart upload tasks are listed in lexicographical order by object name involved in the tasks.</p>
<div class="note" id="note14528254193716"><a name="note14528254193716"></a><a name="note14528254193716"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p252885413370"><a name="p252885413370"></a><a name="p252885413370"></a>For details about how to use this parameter, see <a href="examples-of-listing-multipart-upload-tasks.md">Examples of Listing Multipart Upload Tasks</a>.</p>
</div></div>
</td>
</tr>
<tr id="row6215919119"><td class="cellrowborder" valign="top" width="18.181818181818183%" headers="mcps1.1.4.1.1 "><p id="p118791410119"><a name="p118791410119"></a><a name="p118791410119"></a>uploadIdMarker</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p1189914161120"><a name="p1189914161120"></a><a name="p1189914161120"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="56.565656565656575%" headers="mcps1.1.4.1.3 "><p id="p1190101461115"><a name="p1190101461115"></a><a name="p1190101461115"></a>Indicates the upload ID after which the multipart upload listing begins. This parameter must be used together with <strong id="b11400330113014"><a name="b11400330113014"></a><a name="b11400330113014"></a>marker</strong>. All returned multipart upload tasks are listed in lexicographical order by object name and upload ID involved in the tasks.</p>
</td>
</tr>
<tr id="row1256251542913"><td class="cellrowborder" valign="top" width="18.181818181818183%" headers="mcps1.1.4.1.1 "><p id="p29791221132910"><a name="p29791221132910"></a><a name="p29791221132910"></a>limit</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p119801421162916"><a name="p119801421162916"></a><a name="p119801421162916"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="56.565656565656575%" headers="mcps1.1.4.1.3 "><p id="p1298142132912"><a name="p1298142132912"></a><a name="p1298142132912"></a>Maximum number of objects that can be listed. If the value is less than or equal to <strong id="b13187171310819"><a name="b13187171310819"></a><a name="b13187171310819"></a>0</strong>, all objects are listed.</p>
<div class="note" id="note11900183831617"><a name="note11900183831617"></a><a name="note11900183831617"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p17979111612385"><a name="p17979111612385"></a><a name="p17979111612385"></a>If there are a large number of multipart upload tasks in a bucket, you are advised to set this parameter to limit the number of multipart upload tasks each time. If not all tasks are listed, <strong id="b264333592411"><a name="b264333592411"></a><a name="b264333592411"></a>marker</strong> and <strong id="b1694848122418"><a name="b1694848122418"></a><a name="b1694848122418"></a>uploadIdMarker</strong> of the next request will be returned in the result, which you can use to list the remaining tasks.</p>
</div></div>
</td>
</tr>
<tr id="row1643849183617"><td class="cellrowborder" valign="top" width="18.181818181818183%" headers="mcps1.1.4.1.1 "><p id="p153951131317"><a name="p153951131317"></a><a name="p153951131317"></a>config</p>
</td>
<td class="cellrowborder" valign="top" width="25.252525252525253%" headers="mcps1.1.4.1.2 "><p id="p12395135316"><a name="p12395135316"></a><a name="p12395135316"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="56.565656565656575%" headers="mcps1.1.4.1.3 "><p id="p43952034313"><a name="p43952034313"></a><a name="p43952034313"></a>User-defined configuration file for executing a command. For details about parameters that can be configured, see <a href="parameter-description.md">Parameter Description</a>.</p>
</td>
</tr>
</tbody>
</table>

## Running Example<a name="section17470142316227"></a>

-   Take the Windows OS as an example. Run the  **obsutil ls obs://bucket-test -m -limit=10**  command to query the multipart upload tasks in the bucket.

```
obsutil ls obs://bucket-test -m -limit=10

Listing multipart uploads.

Upload list:
Key                                               Initiated
StorageClass        UploadId
obs://bucket-test/aaa                                  2018-11-27T03:49:07Z
standard            000001675348ED21860C3F61EF955BD3

obs://bucket-test/dir1/10GB.txt                        2018-11-07T06:58:09Z
standard            00000166ECF6CF7C860D1DBAF3F76013

obs://bucket-test/dir1/1GB.txt                         2018-11-07T06:58:09Z
standard            00000166ECF6CF6F860B7FBE95D01B03

obs://bucket-test/dir1/50GB.txt                        2018-11-07T06:58:09Z
standard            00000166ECF6CF86860D1DC2C8E8F66B

obs://bucket-test/dir1/5GB.txt                         2018-11-07T06:58:09Z
standard            00000166ECF6CF75860CDA7780CB52C3

obs://bucket-test/test11/20GB.txt                      2018-11-27T08:21:26Z
standard            0000016754423D24860CA8A4D06C2054

Folder number is: 0
Upload number is: 6
```

-   For more examples, see  [Examples of Listing Multipart Upload Tasks](examples-of-listing-multipart-upload-tasks.md).

