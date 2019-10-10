# Querying Bucket Properties<a name="EN-US_TOPIC_0141181364"></a>

## Function<a name="section1479112110815"></a>

You can use this command to query the basic properties of a bucket, including its default storage class, region, version ID, storage usage, bucket quota, and the number of objects in the bucket.

## Command Line Structure<a name="section1220752192216"></a>

-   In Windows

    ```
    obsutil stat obs://bucket [-acl] [-config=xxx]
    ```

-   In Linux or macOS

    ```
    ./obsutil stat obs://bucket [-acl] [-config=xxx]
    ```


## Parameter Description<a name="section6559191102418"></a>

<a name="table10831182114445"></a>
<table><thead align="left"><tr id="row683212154419"><th class="cellrowborder" valign="top" width="17%" id="mcps1.1.4.1.1"><p id="p118329219446"><a name="p118329219446"></a><a name="p118329219446"></a>Parameter</p>
</th>
<th class="cellrowborder" valign="top" width="25%" id="mcps1.1.4.1.2"><p id="p15137125919108"><a name="p15137125919108"></a><a name="p15137125919108"></a>Optional or Mandatory</p>
</th>
<th class="cellrowborder" valign="top" width="57.99999999999999%" id="mcps1.1.4.1.3"><p id="p12832121184414"><a name="p12832121184414"></a><a name="p12832121184414"></a>Description</p>
</th>
</tr>
</thead>
<tbody><tr id="row108328217449"><td class="cellrowborder" valign="top" width="17%" headers="mcps1.1.4.1.1 "><p id="p64495172515"><a name="p64495172515"></a><a name="p64495172515"></a>bucket</p>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.1.4.1.2 "><p id="p154316502519"><a name="p154316502519"></a><a name="p154316502519"></a>Mandatory</p>
</td>
<td class="cellrowborder" valign="top" width="57.99999999999999%" headers="mcps1.1.4.1.3 "><p id="p17425512259"><a name="p17425512259"></a><a name="p17425512259"></a>Bucket name</p>
</td>
</tr>
<tr id="row141411953161116"><td class="cellrowborder" valign="top" width="17%" headers="mcps1.1.4.1.1 "><p id="p814111533115"><a name="p814111533115"></a><a name="p814111533115"></a>acl</p>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.1.4.1.2 "><p id="p181412534110"><a name="p181412534110"></a><a name="p181412534110"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="57.99999999999999%" headers="mcps1.1.4.1.3 "><p id="p20141165311112"><a name="p20141165311112"></a><a name="p20141165311112"></a>Queries the access control policies of the bucket while querying bucket properties.</p>
</td>
</tr>
<tr id="row10113026103410"><td class="cellrowborder" valign="top" width="17%" headers="mcps1.1.4.1.1 "><p id="p153951131317"><a name="p153951131317"></a><a name="p153951131317"></a>config</p>
</td>
<td class="cellrowborder" valign="top" width="25%" headers="mcps1.1.4.1.2 "><p id="p12395135316"><a name="p12395135316"></a><a name="p12395135316"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="57.99999999999999%" headers="mcps1.1.4.1.3 "><p id="p43952034313"><a name="p43952034313"></a><a name="p43952034313"></a>User-defined configuration file for executing a command. For details about parameters that can be configured, see <a href="parameter-description.md">Parameter Description</a>.</p>
</td>
</tr>
</tbody>
</table>

## Response<a name="section6926520122416"></a>

<a name="table992610203244"></a>
<table><thead align="left"><tr id="row892913208248"><th class="cellrowborder" valign="top" width="21.67%" id="mcps1.1.3.1.1"><p id="p1392992019245"><a name="p1392992019245"></a><a name="p1392992019245"></a>Field</p>
</th>
<th class="cellrowborder" valign="top" width="78.33%" id="mcps1.1.3.1.2"><p id="p19318207246"><a name="p19318207246"></a><a name="p19318207246"></a>Description</p>
</th>
</tr>
</thead>
<tbody><tr id="row79320208248"><td class="cellrowborder" valign="top" width="21.67%" headers="mcps1.1.3.1.1 "><p id="p1093217203246"><a name="p1093217203246"></a><a name="p1093217203246"></a>Bucket</p>
</td>
<td class="cellrowborder" valign="top" width="78.33%" headers="mcps1.1.3.1.2 "><p id="p169337207245"><a name="p169337207245"></a><a name="p169337207245"></a>Bucket name</p>
</td>
</tr>
<tr id="row1753615812247"><td class="cellrowborder" valign="top" width="21.67%" headers="mcps1.1.3.1.1 "><p id="p15536155820245"><a name="p15536155820245"></a><a name="p15536155820245"></a>StorageClass</p>
</td>
<td class="cellrowborder" valign="top" width="78.33%" headers="mcps1.1.3.1.2 "><p id="p3536758162410"><a name="p3536758162410"></a><a name="p3536758162410"></a>Default storage class of the bucket</p>
</td>
</tr>
<tr id="row98643152253"><td class="cellrowborder" valign="top" width="21.67%" headers="mcps1.1.3.1.1 "><p id="p8864415102517"><a name="p8864415102517"></a><a name="p8864415102517"></a>Location</p>
</td>
<td class="cellrowborder" valign="top" width="78.33%" headers="mcps1.1.3.1.2 "><p id="p1886416157252"><a name="p1886416157252"></a><a name="p1886416157252"></a>Region where the bucket resides</p>
</td>
</tr>
<tr id="row116921551192713"><td class="cellrowborder" valign="top" width="21.67%" headers="mcps1.1.3.1.1 "><p id="p0693195172712"><a name="p0693195172712"></a><a name="p0693195172712"></a>ObsVersion</p>
</td>
<td class="cellrowborder" valign="top" width="78.33%" headers="mcps1.1.3.1.2 "><p id="p1269335113275"><a name="p1269335113275"></a><a name="p1269335113275"></a>Version of the bucket</p>
</td>
</tr>
<tr id="row5233151605412"><td class="cellrowborder" valign="top" width="21.67%" headers="mcps1.1.3.1.1 "><p id="p323314166548"><a name="p323314166548"></a><a name="p323314166548"></a>BucketType</p>
</td>
<td class="cellrowborder" valign="top" width="78.33%" headers="mcps1.1.3.1.2 "><p id="p16233171615541"><a name="p16233171615541"></a><a name="p16233171615541"></a>Type of a bucket. <strong id="b12748202742519"><a name="b12748202742519"></a><a name="b12748202742519"></a>OBJECT</strong> indicates a bucket for object storage. <strong id="b3351244152512"><a name="b3351244152512"></a><a name="b3351244152512"></a>POSIX</strong> indicates a bucket used as a parallel file system.</p>
</td>
</tr>
<tr id="row14992172352513"><td class="cellrowborder" valign="top" width="21.67%" headers="mcps1.1.3.1.1 "><p id="p69926234252"><a name="p69926234252"></a><a name="p69926234252"></a>ObjectNumber</p>
</td>
<td class="cellrowborder" valign="top" width="78.33%" headers="mcps1.1.3.1.2 "><p id="p46322324"><a name="p46322324"></a><a name="p46322324"></a>Number of objects in the bucket</p>
</td>
</tr>
<tr id="row1272082862513"><td class="cellrowborder" valign="top" width="21.67%" headers="mcps1.1.3.1.1 "><p id="p4720528142516"><a name="p4720528142516"></a><a name="p4720528142516"></a>Size</p>
</td>
<td class="cellrowborder" valign="top" width="78.33%" headers="mcps1.1.3.1.2 "><p id="p872012286252"><a name="p872012286252"></a><a name="p872012286252"></a>Storage usage of the bucket, in bytes</p>
</td>
</tr>
<tr id="row48453315255"><td class="cellrowborder" valign="top" width="21.67%" headers="mcps1.1.3.1.1 "><p id="p14845531172510"><a name="p14845531172510"></a><a name="p14845531172510"></a>Quota</p>
</td>
<td class="cellrowborder" valign="top" width="78.33%" headers="mcps1.1.3.1.2 "><p id="p11845131182516"><a name="p11845131182516"></a><a name="p11845131182516"></a>Bucket quota. Value <strong id="b29639372521"><a name="b29639372521"></a><a name="b29639372521"></a>0</strong> indicates that no upper limit is set for the bucket quota.</p>
</td>
</tr>
<tr id="row16285968588"><td class="cellrowborder" valign="top" width="21.67%" headers="mcps1.1.3.1.1 "><p id="p5286146135817"><a name="p5286146135817"></a><a name="p5286146135817"></a>Acl</p>
</td>
<td class="cellrowborder" valign="top" width="78.33%" headers="mcps1.1.3.1.2 "><p id="p182864618582"><a name="p182864618582"></a><a name="p182864618582"></a>Access control policy of the bucket</p>
</td>
</tr>
</tbody>
</table>

## Running Example<a name="section15899161919244"></a>

-   Take the Windows OS as an example. Run the  **obsutil stat obs://bucket-test**  command to query the basic properties of bucket  **bucket-test**.

```
obsutil stat obs://bucket-test

Bucket:
  obs://bucket-test
StorageClass:
  standard
Location:
  southchina
ObsVersion:
  3.0
ObjectNumber:
  8005
Size:
  320076506
Quota:
  0
```

