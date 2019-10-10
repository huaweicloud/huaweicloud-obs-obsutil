# Listing Objects<a name="EN-US_TOPIC_0142009729"></a>

## Function<a name="section1479112110815"></a>

You can use this command to query objects or object versions in a bucket. All objects are listed in lexicographical order by object name and version ID.

## Command Line Structure<a name="section1220752192216"></a>

-   In Windows

    ```
    obsutil ls obs://bucket[/prefix] [-s] [-d] [-v] [-marker=xxx] [-versionIdMarker=xxx] [-bf=xxx] [-limit=1] [-config=xxx]
    ```

-   In Linux or macOS

    ```
    ./obsutil ls obs://bucket[/prefix] [-s] [-d] [-v] [-marker=xxx] [-versionIdMarker=xxx] [-bf=xxx] [-limit=1] [-config=xxx]
    ```


## Parameter Description<a name="section6559191102418"></a>

<a name="table10831182114445"></a>
<table><thead align="left"><tr id="row683212154419"><th class="cellrowborder" valign="top" width="18.19%" id="mcps1.1.4.1.1"><p id="p118329219446"><a name="p118329219446"></a><a name="p118329219446"></a>Parameter</p>
</th>
<th class="cellrowborder" valign="top" width="26.16%" id="mcps1.1.4.1.2"><p id="p15137125919108"><a name="p15137125919108"></a><a name="p15137125919108"></a>Optional or Mandatory</p>
</th>
<th class="cellrowborder" valign="top" width="55.65%" id="mcps1.1.4.1.3"><p id="p12832121184414"><a name="p12832121184414"></a><a name="p12832121184414"></a>Description</p>
</th>
</tr>
</thead>
<tbody><tr id="row108328217449"><td class="cellrowborder" valign="top" width="18.19%" headers="mcps1.1.4.1.1 "><p id="p64495172515"><a name="p64495172515"></a><a name="p64495172515"></a>bucket</p>
</td>
<td class="cellrowborder" valign="top" width="26.16%" headers="mcps1.1.4.1.2 "><p id="p154316502519"><a name="p154316502519"></a><a name="p154316502519"></a>Mandatory</p>
</td>
<td class="cellrowborder" valign="top" width="55.65%" headers="mcps1.1.4.1.3 "><p id="p17425512259"><a name="p17425512259"></a><a name="p17425512259"></a>Bucket name</p>
</td>
</tr>
<tr id="row77261271514"><td class="cellrowborder" valign="top" width="18.19%" headers="mcps1.1.4.1.1 "><p id="p157268710117"><a name="p157268710117"></a><a name="p157268710117"></a>prefix</p>
</td>
<td class="cellrowborder" valign="top" width="26.16%" headers="mcps1.1.4.1.2 "><p id="p19727157912"><a name="p19727157912"></a><a name="p19727157912"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="55.65%" headers="mcps1.1.4.1.3 "><p id="p177272717115"><a name="p177272717115"></a><a name="p177272717115"></a>Prefix of an object name for listing objects</p>
<div class="note" id="note129881917144317"><a name="note129881917144317"></a><a name="note129881917144317"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p2479161820431"><a name="p2479161820431"></a><a name="p2479161820431"></a>If this parameter is left blank, all objects in the bucket are listed.</p>
</div></div>
</td>
</tr>
<tr id="row167881281811"><td class="cellrowborder" valign="top" width="18.19%" headers="mcps1.1.4.1.1 "><p id="p37889281012"><a name="p37889281012"></a><a name="p37889281012"></a>s</p>
</td>
<td class="cellrowborder" valign="top" width="26.16%" headers="mcps1.1.4.1.2 "><p id="p177888281316"><a name="p177888281316"></a><a name="p177888281316"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.65%" headers="mcps1.1.4.1.3 "><p id="p1378816281111"><a name="p1378816281111"></a><a name="p1378816281111"></a>Displays simplified query result.</p>
<div class="note" id="note1452085114212"><a name="note1452085114212"></a><a name="note1452085114212"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p252035122111"><a name="p252035122111"></a><a name="p252035122111"></a>In the simplified format, the returned result contains only the object name.</p>
</div></div>
</td>
</tr>
<tr id="row105292037563"><td class="cellrowborder" valign="top" width="18.19%" headers="mcps1.1.4.1.1 "><p id="p145294376611"><a name="p145294376611"></a><a name="p145294376611"></a>d</p>
</td>
<td class="cellrowborder" valign="top" width="26.16%" headers="mcps1.1.4.1.2 "><p id="p18226142117716"><a name="p18226142117716"></a><a name="p18226142117716"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.65%" headers="mcps1.1.4.1.3 "><p id="p853014371767"><a name="p853014371767"></a><a name="p853014371767"></a>Lists only objects and subdirectories in the current directory, instead of recursively listing all objects and subdirectories.</p>
<div class="note" id="note1255192611160"><a name="note1255192611160"></a><a name="note1255192611160"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p156102661616"><a name="p156102661616"></a><a name="p156102661616"></a>According to the naming conventions in OBS, a slash (/) is used as the directory separator.</p>
</div></div>
</td>
</tr>
<tr id="row68204468610"><td class="cellrowborder" valign="top" width="18.19%" headers="mcps1.1.4.1.1 "><p id="p3820124620619"><a name="p3820124620619"></a><a name="p3820124620619"></a>v</p>
</td>
<td class="cellrowborder" valign="top" width="26.16%" headers="mcps1.1.4.1.2 "><p id="p982013461162"><a name="p982013461162"></a><a name="p982013461162"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.65%" headers="mcps1.1.4.1.3 "><p id="p108200466610"><a name="p108200466610"></a><a name="p108200466610"></a>Lists versions of an object in a bucket. The result contains the latest version and historical versions (if any) of the object.</p>
</td>
</tr>
<tr id="row182013491164"><td class="cellrowborder" valign="top" width="18.19%" headers="mcps1.1.4.1.1 "><p id="p1201174920618"><a name="p1201174920618"></a><a name="p1201174920618"></a>marker</p>
</td>
<td class="cellrowborder" valign="top" width="26.16%" headers="mcps1.1.4.1.2 "><p id="p1020113491669"><a name="p1020113491669"></a><a name="p1020113491669"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.65%" headers="mcps1.1.4.1.3 "><p id="p820104913610"><a name="p820104913610"></a><a name="p820104913610"></a>Object name to start with when listing objects in a bucket. All objects are listed in lexicographical order by object name.</p>
<div class="note" id="note14528254193716"><a name="note14528254193716"></a><a name="note14528254193716"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p252885413370"><a name="p252885413370"></a><a name="p252885413370"></a>For details about how to use this parameter, see <a href="list-examples.md">List Examples</a>.</p>
</div></div>
</td>
</tr>
<tr id="row1558665118612"><td class="cellrowborder" valign="top" width="18.19%" headers="mcps1.1.4.1.1 "><p id="p858665116612"><a name="p858665116612"></a><a name="p858665116612"></a>versionIdMarker</p>
</td>
<td class="cellrowborder" valign="top" width="26.16%" headers="mcps1.1.4.1.2 "><p id="p16675153711112"><a name="p16675153711112"></a><a name="p16675153711112"></a>Optional (additional parameter). It must be used together with the <strong id="b18703701119"><a name="b18703701119"></a><a name="b18703701119"></a>v</strong> and <strong id="b9569143121112"><a name="b9569143121112"></a><a name="b9569143121112"></a>marker</strong> parameters.</p>
</td>
<td class="cellrowborder" valign="top" width="55.65%" headers="mcps1.1.4.1.3 "><p id="p95866511617"><a name="p95866511617"></a><a name="p95866511617"></a>Version ID to start with when listing versions of objects in a bucket. All versions and objects are listed in lexicographical order by object name and version ID.</p>
<div class="note" id="note183041831558"><a name="note183041831558"></a><a name="note183041831558"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p93043312514"><a name="p93043312514"></a><a name="p93043312514"></a>If the value of <strong id="b842352706143726"><a name="b842352706143726"></a><a name="b842352706143726"></a>versionIdMarker</strong> is not a version ID specified by <strong id="b842352706143743"><a name="b842352706143743"></a><a name="b842352706143743"></a>marker</strong>, <strong id="b1781920833143750"><a name="b1781920833143750"></a><a name="b1781920833143750"></a>versionIdMarker</strong> is invalid.</p>
</div></div>
</td>
</tr>
<tr id="row193001710154315"><td class="cellrowborder" valign="top" width="18.19%" headers="mcps1.1.4.1.1 "><p id="p63016107437"><a name="p63016107437"></a><a name="p63016107437"></a>bf</p>
</td>
<td class="cellrowborder" valign="top" width="26.16%" headers="mcps1.1.4.1.2 "><p id="p10301510184317"><a name="p10301510184317"></a><a name="p10301510184317"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.65%" headers="mcps1.1.4.1.3 "><p id="p1930171012438"><a name="p1930171012438"></a><a name="p1930171012438"></a>Display formats of bytes in the listing result. Possible values are:</p>
<a name="ul1273864144412"></a><a name="ul1273864144412"></a><ul id="ul1273864144412"><li>human-readable</li><li>raw</li></ul>
<div class="note" id="note11602143394810"><a name="note11602143394810"></a><a name="note11602143394810"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p6603633154818"><a name="p6603633154818"></a><a name="p6603633154818"></a>If this parameter is not configured, the display format of bytes in the result is determined by the <strong id="b12496171034012"><a name="b12496171034012"></a><a name="b12496171034012"></a>humanReadableFormat</strong> parameter in the configuration file.</p>
</div></div>
</td>
</tr>
<tr id="row23091241125114"><td class="cellrowborder" valign="top" width="18.19%" headers="mcps1.1.4.1.1 "><p id="p1918111241654"><a name="p1918111241654"></a><a name="p1918111241654"></a>limit</p>
</td>
<td class="cellrowborder" valign="top" width="26.16%" headers="mcps1.1.4.1.2 "><p id="p218113247513"><a name="p218113247513"></a><a name="p218113247513"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.65%" headers="mcps1.1.4.1.3 "><p id="p310214615616"><a name="p310214615616"></a><a name="p310214615616"></a>Maximum number of objects that can be listed. If the value is less than or equal to 0, all objects are listed. If it is left blank, 1000 objects are listed by default.</p>
<div class="note" id="note1997814166384"><a name="note1997814166384"></a><a name="note1997814166384"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p17979111612385"><a name="p17979111612385"></a><a name="p17979111612385"></a>If there are a large number of objects in a bucket, you are advised to set this parameter to limit the number of objects to be listed each time. If not all objects are listed, <strong id="b1440343663211"><a name="b1440343663211"></a><a name="b1440343663211"></a>marker</strong> and <strong id="b350211407322"><a name="b350211407322"></a><a name="b350211407322"></a>versionIdMarker</strong> of the next request will be returned in the result, which you can use to list the remaining objects.</p>
</div></div>
</td>
</tr>
<tr id="row1095422313515"><td class="cellrowborder" valign="top" width="18.19%" headers="mcps1.1.4.1.1 "><p id="p153951131317"><a name="p153951131317"></a><a name="p153951131317"></a>config</p>
</td>
<td class="cellrowborder" valign="top" width="26.16%" headers="mcps1.1.4.1.2 "><p id="p12395135316"><a name="p12395135316"></a><a name="p12395135316"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.65%" headers="mcps1.1.4.1.3 "><p id="p43952034313"><a name="p43952034313"></a><a name="p43952034313"></a>User-defined configuration file for executing a command. For details about parameters that can be configured, see <a href="parameter-description.md">Parameter Description</a>.</p>
</td>
</tr>
</tbody>
</table>

## Running Example<a name="section15899161919244"></a>

-   Take the Windows OS as an example. Run the  **obsutil ls obs://bucket-test -limit=10**  command to list objects in the bucket.

```
obsutil ls obs://bucket-test -limit=10

Folder list:
obs://bucket-test/api/

Object list:
key                                               LastModified                  Size      StorageClass        ETag                
obs://bucket-test/AUTHORS                         2018-11-16T02:15:49Z          33243     standard            "796393c1eaf502ef56a85c2ceb640aea"

obs://bucket-test/CONTRIBUTING.md                 2018-11-16T02:15:49Z          1366      standard            "12d93325ba6131f852daecd18dd65edc"

obs://bucket-test/CONTRIBUTORS                    2018-11-16T02:15:49Z          45710     standard            "b486b5003e6215c9199e86ab3ccec9fa"

obs://bucket-test/LICENSE                         2018-11-16T02:15:49Z          1479      standard            "5d4950ecb7b26d2c5e4e7b4e0dd74707"

obs://bucket-test/PATENTS                         2018-11-16T02:15:49Z          1303      standard            "3a55d95595a6f9e37dee53826b4daff2"

obs://bucket-test/README.md                       2018-11-16T02:15:49Z          1399      standard            "97351fd7946b9ea021a31a86ba2a10ab"

obs://bucket-test/VERSION                         2018-11-16T02:15:49Z          7         standard            "43d93b553855b0e1fc67e31c28c07b65"

obs://bucket-test/api/README                      2018-11-16T02:15:49Z          521       standard            "4e9e63a87075df60cdf65c8ce9e92117"

obs://bucket-test/api/except.txt                  2018-11-16T02:15:49Z          20194     standard            "8eb96de3f60447e2f09a7531c99fb3ee"

Next marker is: api/except.txt
Folder number is: 1
File number is: 9
```

-   For more examples, see  [List Examples](list-examples.md).

