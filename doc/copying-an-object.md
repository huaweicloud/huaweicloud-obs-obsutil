# Copying an Object<a name="EN-US_TOPIC_0142359355"></a>

## Function<a name="section1479112110815"></a>

You can use this command to copy a single object or copy objects in batches by a specified object name prefix.

>![](public_sys-resources/icon-notice.gif) **NOTICE:**   
>-   Do not change the source objects in the OBS bucket when copying a single object or objects in batches. Otherwise, the operation may fail or data may be inconsistent.  
>-   If the storage class of the object to be copied is  **cold**, you must restore the object to be copied first. Otherwise, the copy fails.  
>-   To copy objects, you must have the read permission on the objects to be copied and the write permission on the destination bucket.  
>-   If the client-side cross-region replication function is not enabled, ensure that the source bucket and destination bucket are in the same region.  

## Command Line Structure<a name="section49408320267"></a>

-   In Windows
    -   Copying a single object

        ```
        obsutil cp obs://srcbucket/key obs://dstbucket/[dest] [-dryRun][-u] [-crr] [-vlength] [-vmd5] [-p=1] [-threshold=52428800] [-versionId=xxx] [-acl=xxx] [-sc=xxx] [-meta=aaa:bbb#ccc:ddd] [-ps=auto] [-cpd=xxx] [-fr] [-o=xxx] [-config=xxx]
        ```

    -   Copying objects in batches

        ```
        obsutil cp obs://srcbucket[/key] obs://dstbucket[/dest] -r [-dryRun][-f] [-flat] [-u] [-crr] [-vlength] [-vmd5] [-j=1] [-p=1] [-threshold=52428800] [-acl=xxx] [-sc=xxx] [-meta=aaa:bbb#ccc:ddd] [-ps=auto] [-include=*.xxx] [-exclude=*.xxx] [-timeRange=time1-time2] [-mf] [-o=xxx] [-cpd=xxx] [-config=xxx]
        ```


-   In Linux or macOS
    -   Copying a single object

        ```
        ./obsutil cp obs://srcbucket/key obs://dstbucket/[dest] [-dryRun] [-u] [-crr] [-vlength] [-vmd5] [-p=1] [-threshold=52428800] [-versionId=xxx] [-acl=xxx] [-sc=xxx] [-meta=aaa:bbb#ccc:ddd] [-ps=auto] [-cpd=xxx] [-fr] [-o=xxx] [-config=xxx]
        ```

    -   Copying objects in batches

        ```
        ./obsutil cp obs://srcbucket[/key] obs://dstbucket[/dest] -r [-dryRun] [-f] [-flat] [-u] [-crr] [-vlength] [-vmd5] [-j=1] [-p=1] [-threshold=52428800] [-acl=xxx] [-sc=xxx] [-meta=aaa:bbb#ccc:ddd] [-ps=auto] [-include=*.xxx] [-exclude=*.xxx] [-timeRange=time1-time2] [-mf] [-o=xxx] [-cpd=xxx] [-config=xxx]
        ```



>![](public_sys-resources/icon-note.gif) **NOTE:**   
>-   The source path and destination path cannot be the same.  
>-   The source path and destination path cannot be partly overlapped either. If the source path overlaps with the prefix of the destination path, recursive replication applies. If the destination path overlaps with the prefix of the source path, the replication may overwrite objects in the source path.  

## Parameter Description<a name="section3940203142610"></a>

<a name="table125609307312"></a>
<table><thead align="left"><tr id="row456603043117"><th class="cellrowborder" valign="top" width="17.57%" id="mcps1.1.4.1.1"><p id="p85680303316"><a name="p85680303316"></a><a name="p85680303316"></a>Parameter</p>
</th>
<th class="cellrowborder" valign="top" width="18.529999999999998%" id="mcps1.1.4.1.2"><p id="p14572130153112"><a name="p14572130153112"></a><a name="p14572130153112"></a>Optional or Mandatory</p>
</th>
<th class="cellrowborder" valign="top" width="63.9%" id="mcps1.1.4.1.3"><p id="p15751930133118"><a name="p15751930133118"></a><a name="p15751930133118"></a>Description</p>
</th>
</tr>
</thead>
<tbody><tr id="row20577630203112"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p18579630143115"><a name="p18579630143115"></a><a name="p18579630143115"></a>srcbucket</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p458314302313"><a name="p458314302313"></a><a name="p458314302313"></a>Mandatory</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p34544306544"><a name="p34544306544"></a><a name="p34544306544"></a>Source bucket name</p>
</td>
</tr>
<tr id="row1059316307314"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p159616306318"><a name="p159616306318"></a><a name="p159616306318"></a>dstbucket</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p1559933073119"><a name="p1559933073119"></a><a name="p1559933073119"></a>Mandatory</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p95994308315"><a name="p95994308315"></a><a name="p95994308315"></a>Destination bucket name</p>
</td>
</tr>
<tr id="row169347719447"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p1811911004419"><a name="p1811911004419"></a><a name="p1811911004419"></a>dest</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p9123510204415"><a name="p9123510204415"></a><a name="p9123510204415"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p1839313263477"><a name="p1839313263477"></a><a name="p1839313263477"></a>Indicates the destination object name when copying an object, or the name prefix of destination objects when copying objects in batches.</p>
</td>
</tr>
<tr id="row106007303311"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p1060363010318"><a name="p1060363010318"></a><a name="p1060363010318"></a>key</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p8535551165513"><a name="p8535551165513"></a><a name="p8535551165513"></a>Mandatory for copying an object.</p>
<p id="p7265911561"><a name="p7265911561"></a><a name="p7265911561"></a>Optional for copying objects in batches.</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p176052030183114"><a name="p176052030183114"></a><a name="p176052030183114"></a>Indicates the source object name when copying an object, or the name prefix of source objects when copying objects in batches.</p>
<p id="p36061309319"><a name="p36061309319"></a><a name="p36061309319"></a>The rules are as follows:</p>
<a name="ul1560743053116"></a><a name="ul1560743053116"></a><ul id="ul1560743053116"><li>This parameter cannot be left blank when copying an object. If <strong id="b17382225533"><a name="b17382225533"></a><a name="b17382225533"></a>dest</strong> is left blank, the source object is copied to the root directory of the destination bucket. If the value of <strong id="b838316257318"><a name="b838316257318"></a><a name="b838316257318"></a>dest</strong> ends with a slash (/), the destination object name is the value of <strong id="b238462516311"><a name="b238462516311"></a><a name="b238462516311"></a>dest</strong> plus the source object name. Otherwise, the destination object name is the value of <strong id="b14384325336"><a name="b14384325336"></a><a name="b14384325336"></a>dest</strong>.</li><li>If this parameter is left blank when copying objects in batches, all objects in the source bucket are copied. If not, objects whose name prefix is the set value in the source bucket are copied. The rules for confirming the name of the destination object are as follows:<a name="ul101911711114619"></a><a name="ul101911711114619"></a><ul id="ul101911711114619"><li>If the value of <strong id="b175610418258"><a name="b175610418258"></a><a name="b175610418258"></a>dest</strong> ends with a slash (/), the destination object name is the value of <strong id="b1756312281499"><a name="b1756312281499"></a><a name="b1756312281499"></a>dest</strong> plus the source object name.</li><li>If the value of <strong id="b198673121165"><a name="b198673121165"></a><a name="b198673121165"></a>dest</strong> does not end with a slash (/), the destination object name is <em id="i1144918511760"><a name="i1144918511760"></a><a name="i1144918511760"></a>the value of </em><em id="i651522503618"><a name="i651522503618"></a><a name="i651522503618"></a><strong id="b92851281163"><a name="b92851281163"></a><a name="b92851281163"></a>dest</strong></em><strong id="b10281031178"><a name="b10281031178"></a><a name="b10281031178"></a>/</strong><em id="i563313811717"><a name="i563313811717"></a><a name="i563313811717"></a>source object name</em>.</li></ul>
</li></ul>
<div class="note" id="note14528254193716"><a name="note14528254193716"></a><a name="note14528254193716"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul17203733184419"></a><a name="ul17203733184419"></a><ul id="ul17203733184419"><li>If this parameter is configured but the <strong id="b882713482518"><a name="b882713482518"></a><a name="b882713482518"></a>flat</strong> parameter is not when copying objects in batches, the name of the source object contains the name prefix of the parent object. If <strong id="b168287489512"><a name="b168287489512"></a><a name="b168287489512"></a>flat</strong> is configured, then the name of the source object does not contain the name prefix of the parent object.</li><li>For details about how to use this parameter, see <a href="copy-examples.md">Copy Examples</a>.</li></ul>
</div></div>
</td>
</tr>
<tr id="row64161323113"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p123051451513"><a name="p123051451513"></a><a name="p123051451513"></a>fr</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p113071451811"><a name="p113071451811"></a><a name="p113071451811"></a>Optional for copying an object (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p73101351816"><a name="p73101351816"></a><a name="p73101351816"></a>Generates an operation result list when copying an object.</p>
</td>
</tr>
<tr id="row19803114816115"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p13236151151117"><a name="p13236151151117"></a><a name="p13236151151117"></a>flat</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p142391351111116"><a name="p142391351111116"></a><a name="p142391351111116"></a>Optional for copying objects in batches (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p624025171110"><a name="p624025171110"></a><a name="p624025171110"></a>The name prefix of the parent object is excluded when copying objects in batches.</p>
</td>
</tr>
<tr id="row1513173918563"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p1537517155414"><a name="p1537517155414"></a><a name="p1537517155414"></a>dryRun</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p13376216548"><a name="p13376216548"></a><a name="p13376216548"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p1137981175411"><a name="p1137981175411"></a><a name="p1137981175411"></a>Conducts a dry run.</p>
</td>
</tr>
<tr id="row4959122013591"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p19960122065918"><a name="p19960122065918"></a><a name="p19960122065918"></a>crr</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p696012075919"><a name="p696012075919"></a><a name="p696012075919"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p1296062012595"><a name="p1296062012595"></a><a name="p1296062012595"></a>Enables the client-side cross-region replication function. In this mode, data is directly copied to the destination bucket from the source bucket through data stream. The buckets can by any two OBS buckets.</p>
<div class="note" id="note4424164104813"><a name="note4424164104813"></a><a name="note4424164104813"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul743879115112"></a><a name="ul743879115112"></a><ul id="ul743879115112"><li>If this parameter is configured, ensure that the configuration of client-side cross-region replication is updated in the configuration file. For details, see <a href="updating-a-configuration-file.md">Updating a Configuration File</a>.</li><li>The configurations of the source bucket and destination bucket are respectively <strong id="b1531810395405"><a name="b1531810395405"></a><a name="b1531810395405"></a>akCrr/skCrr/tokenCrr/endpointCrr</strong> and <strong id="b18642172774116"><a name="b18642172774116"></a><a name="b18642172774116"></a>ak/sk/token/endpoint</strong> in the configuration file.</li></ul>
</div></div>
<div class="notice" id="note19922116154811"><a name="note19922116154811"></a><a name="note19922116154811"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><p id="p4922106154811"><a name="p4922106154811"></a><a name="p4922106154811"></a>After this function is enabled, both upload and download bandwidth are occupied.</p>
</div></div>
</td>
</tr>
<tr id="row145913418207"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p1605745105311"><a name="p1605745105311"></a><a name="p1605745105311"></a>vlength</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p12605114512534"><a name="p12605114512534"></a><a name="p12605114512534"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p1660724595316"><a name="p1660724595316"></a><a name="p1660724595316"></a>Verifies whether the object size in the destination bucket is the same as that in the source bucket after the copy task completes.</p>
<div class="note" id="note57671113172915"><a name="note57671113172915"></a><a name="note57671113172915"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p1768313182910"><a name="p1768313182910"></a><a name="p1768313182910"></a>This parameter must be used together with <strong id="b19902654113813"><a name="b19902654113813"></a><a name="b19902654113813"></a>crr</strong>.</p>
</div></div>
</td>
</tr>
<tr id="row16347124510200"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p136091945105315"><a name="p136091945105315"></a><a name="p136091945105315"></a>vmd5</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p461011452539"><a name="p461011452539"></a><a name="p461011452539"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p19611745185310"><a name="p19611745185310"></a><a name="p19611745185310"></a>Verifies whether the MD5 value of the destination bucket is the same as that of the source bucket after the copy task completes.</p>
<div class="note" id="note13542169102916"><a name="note13542169102916"></a><a name="note13542169102916"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul114913379291"></a><a name="ul114913379291"></a><ul id="ul114913379291"><li>This parameter must be used together with <strong id="b412431819413"><a name="b412431819413"></a><a name="b412431819413"></a>crr</strong>.</li><li>Objects in the source bucket must contain metadata <strong id="b13893135467"><a name="b13893135467"></a><a name="b13893135467"></a>x-obs-md5chksum</strong>. Otherwise, MD5 verification will be skipped.<p id="p1765649132314"><a name="p1765649132314"></a><a name="p1765649132314"></a>After the MD5 value verification is successful, the parameter value is set to the destination object metadata <strong id="b873752561917"><a name="b873752561917"></a><a name="b873752561917"></a>x-obs-md5chksum</strong>, which is used for later MD5 verification during download or copy.</p>
</li></ul>
</div></div>
</td>
</tr>
<tr id="row17612113018317"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p1461418302318"><a name="p1461418302318"></a><a name="p1461418302318"></a>u</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p1615183013110"><a name="p1615183013110"></a><a name="p1615183013110"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p361753019315"><a name="p361753019315"></a><a name="p361753019315"></a>Indicates incremental copy. If this parameter is set, each object can be copied only when it does not exist in the destination bucket, its size is different from the namesake one in the destination bucket, or it has the latest modification time.</p>
</td>
</tr>
<tr id="row3632193093117"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p26341030163117"><a name="p26341030163117"></a><a name="p26341030163117"></a>p</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p9636183093120"><a name="p9636183093120"></a><a name="p9636183093120"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p7638163013112"><a name="p7638163013112"></a><a name="p7638163013112"></a>Indicates the maximum number of concurrent multipart copy tasks when copying an object. The default value is the value of <strong id="b625031412158"><a name="b625031412158"></a><a name="b625031412158"></a>defaultParallels</strong> in the configuration file.</p>
</td>
</tr>
<tr id="row363913014313"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p116401030183113"><a name="p116401030183113"></a><a name="p116401030183113"></a>threshold</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p176431304319"><a name="p176431304319"></a><a name="p176431304319"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p16441830113114"><a name="p16441830113114"></a><a name="p16441830113114"></a>Indicates the threshold for enabling multipart copy, in bytes. The default value is the value of <strong id="b208241652302"><a name="b208241652302"></a><a name="b208241652302"></a>defaultBigfileThreshold</strong> in the configuration file.</p>
<div class="note" id="note1382352674913"><a name="note1382352674913"></a><a name="note1382352674913"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul182059695115"></a><a name="ul182059695115"></a><ul id="ul182059695115"><li>If the size of the object to be copied is smaller than the threshold, copy the object directly. If not, a multipart copy is required.</li><li>If you copy an object directly, no part record is generated, and resumable transmission is not supported.</li><li>This value can contain a capacity unit. For example, <strong id="b20463171820458"><a name="b20463171820458"></a><a name="b20463171820458"></a>1 MB</strong> indicates 1048576 bytes.</li></ul>
</div></div>
</td>
</tr>
<tr id="row145994924514"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p95911457154514"><a name="p95911457154514"></a><a name="p95911457154514"></a>versionId</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p0594135774510"><a name="p0594135774510"></a><a name="p0594135774510"></a>Optional for copying an object (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p1859655744514"><a name="p1859655744514"></a><a name="p1859655744514"></a>Source object version ID that can be specified when copying an object</p>
</td>
</tr>
<tr id="row264453012315"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p964614304315"><a name="p964614304315"></a><a name="p964614304315"></a>acl</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p17648193011318"><a name="p17648193011318"></a><a name="p17648193011318"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p16650163083117"><a name="p16650163083117"></a><a name="p16650163083117"></a>Access control policies for destination objects that can be specified when copying objects. Possible values are:</p>
<a name="ul0651183053115"></a><a name="ul0651183053115"></a><ul id="ul0651183053115"><li>private</li><li>public-read</li><li>public-read-write</li><li>bucket-owner-full-control</li></ul>
<div class="note" id="note1790113183525"><a name="note1790113183525"></a><a name="note1790113183525"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p92982882916"><a name="p92982882916"></a><a name="p92982882916"></a>The preceding four values indicate private read and write, public read, public read and write, and bucket owner full control.</p>
</div></div>
</td>
</tr>
<tr id="row62441623105116"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p13466172515511"><a name="p13466172515511"></a><a name="p13466172515511"></a>sc</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p134681525175117"><a name="p134681525175117"></a><a name="p134681525175117"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p04731025135114"><a name="p04731025135114"></a><a name="p04731025135114"></a>Storage classes of the destination objects that can be specified when copying objects. Possible values are:</p>
<a name="ul175651814214"></a><a name="ul175651814214"></a><ul id="ul175651814214"><li><strong id="b1014985213713"><a name="b1014985213713"></a><a name="b1014985213713"></a>standard</strong>: OBS Standard, which features low access latency and high throughput, and is applicable to storing frequently accessed data (multiple accesses per month averagely) or data that is smaller than 1 MB</li><li><strong id="b413119117415"><a name="b413119117415"></a><a name="b413119117415"></a>warm</strong>: OBS Warm. It is applicable to storing semi-frequently accessed (less than 12 times a year averagely) data that requires quick response.</li><li><strong id="b583211551835"><a name="b583211551835"></a><a name="b583211551835"></a>cold</strong>: OBS Cold. It is secure, durable, and inexpensive, and applicable to archiving rarely-accessed (once a year averagely) data.</li></ul>
</td>
</tr>
<tr id="row8659203023112"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p1366223093111"><a name="p1366223093111"></a><a name="p1366223093111"></a>meta</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p26631230193118"><a name="p26631230193118"></a><a name="p26631230193118"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p17941525135111"><a name="p17941525135111"></a><a name="p17941525135111"></a>Metadata of destination objects that can be specified when copying objects. The format is <em id="i128366111210"><a name="i128366111210"></a><a name="i128366111210"></a>key1</em><strong id="b185348207220"><a name="b185348207220"></a><a name="b185348207220"></a>:</strong><em id="i3536122010217"><a name="i3536122010217"></a><a name="i3536122010217"></a>value1</em><strong id="b9590255213"><a name="b9590255213"></a><a name="b9590255213"></a>#</strong><em id="i1347618252210"><a name="i1347618252210"></a><a name="i1347618252210"></a>key2</em><strong id="b1271314295214"><a name="b1271314295214"></a><a name="b1271314295214"></a>:</strong><em id="i2060123018216"><a name="i2060123018216"></a><a name="i2060123018216"></a>value2</em><strong id="b5820143310212"><a name="b5820143310212"></a><a name="b5820143310212"></a>#</strong><em id="i2013018349212"><a name="i2013018349212"></a><a name="i2013018349212"></a>key3</em><strong id="b176411374219"><a name="b176411374219"></a><a name="b176411374219"></a>:</strong><em id="i9928163710213"><a name="i9928163710213"></a><a name="i9928163710213"></a>value3</em>.</p>
<div class="note" id="note1606229165119"><a name="note1606229165119"></a><a name="note1606229165119"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p686342013559"><a name="p686342013559"></a><a name="p686342013559"></a>The preceding value indicates that the destination objects in the bucket contain three groups of customized metadata after objects are copied: <strong id="b13158204812358"><a name="b13158204812358"></a><a name="b13158204812358"></a>key1:value1</strong>, <strong id="b101591048143519"><a name="b101591048143519"></a><a name="b101591048143519"></a>key2:value2</strong>, and <strong id="b4161124853515"><a name="b4161124853515"></a><a name="b4161124853515"></a>key3:value3</strong>.</p>
</div></div>
</td>
</tr>
<tr id="row466716303317"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p1366917307314"><a name="p1366917307314"></a><a name="p1366917307314"></a>ps</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p267203023115"><a name="p267203023115"></a><a name="p267203023115"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p1167313023117"><a name="p1167313023117"></a><a name="p1167313023117"></a>Indicates the size of each part in a multipart copy task, in bytes. The value ranges from 100 KB to 5 GB. The default value is the value of <strong id="b12730856183616"><a name="b12730856183616"></a><a name="b12730856183616"></a>defaultPartSize</strong> in the configuration file.</p>
<div class="note" id="note12706112664318"><a name="note12706112664318"></a><a name="note12706112664318"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul17189125814398"></a><a name="ul17189125814398"></a><ul id="ul17189125814398"><li>This value can contain a capacity unit. For example, <strong id="b3870123314110"><a name="b3870123314110"></a><a name="b3870123314110"></a>1 MB</strong> indicates 1048576 bytes.</li><li>The parameter can be set to <strong id="b2066915352417"><a name="b2066915352417"></a><a name="b2066915352417"></a>auto</strong>. In this case, obsutil automatically sets the part size for each multipart task based on the source object size.</li></ul>
</div></div>
</td>
</tr>
<tr id="row1367463093119"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p1967612305318"><a name="p1967612305318"></a><a name="p1967612305318"></a>cpd</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p1367810308316"><a name="p1367810308316"></a><a name="p1367810308316"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p46792307316"><a name="p46792307316"></a><a name="p46792307316"></a>Indicates the folder where the part records reside. The default value is <strong id="b12387185613189"><a name="b12387185613189"></a><a name="b12387185613189"></a>.obsutil_checkpoint</strong>, the subfolder in the home directory of the user who executes obsutil commands.</p>
<div class="note" id="note22601842192718"><a name="note22601842192718"></a><a name="note22601842192718"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p118866577266"><a name="p118866577266"></a><a name="p118866577266"></a>A part record is generated during a multipart copy and saved to the <strong id="b173819156497"><a name="b173819156497"></a><a name="b173819156497"></a>copy</strong> subfolder. After the copy succeeds, its part record is deleted automatically. If the copy fails or is suspended, the system attempts to resume the task according to its part record when you perform the copy the next time.</p>
</div></div>
</td>
</tr>
<tr id="row12681173043113"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p1683230123111"><a name="p1683230123111"></a><a name="p1683230123111"></a>r</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p1368433083110"><a name="p1368433083110"></a><a name="p1368433083110"></a>Mandatory for copying objects in batches (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p12686930183119"><a name="p12686930183119"></a><a name="p12686930183119"></a>Copies objects in batches based on a specified name prefix of objects in the source bucket.</p>
</td>
</tr>
<tr id="row1568763019310"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p2688183018315"><a name="p2688183018315"></a><a name="p2688183018315"></a>f</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p1691193012317"><a name="p1691193012317"></a><a name="p1691193012317"></a>Optional for copying objects in batches (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p1269315302318"><a name="p1269315302318"></a><a name="p1269315302318"></a>Runs in force mode.</p>
</td>
</tr>
<tr id="row20694163013111"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p269613053110"><a name="p269613053110"></a><a name="p269613053110"></a>j</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p170013055218"><a name="p170013055218"></a><a name="p170013055218"></a>Optional for copying objects in batches (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p4699130153112"><a name="p4699130153112"></a><a name="p4699130153112"></a>Indicates the maximum number of concurrent tasks for copying objects in batches. The default value is the value of <strong id="b6802155075017"><a name="b6802155075017"></a><a name="b6802155075017"></a>defaultJobs</strong> in the configuration file.</p>
<div class="note" id="note891964620819"><a name="note891964620819"></a><a name="note891964620819"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p1091964618820"><a name="p1091964618820"></a><a name="p1091964618820"></a>The value is ensured to be greater than or equal to 1.</p>
</div></div>
</td>
</tr>
<tr id="row8334161131517"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p7660121421518"><a name="p7660121421518"></a><a name="p7660121421518"></a>exclude</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p176621714131512"><a name="p176621714131512"></a><a name="p176621714131512"></a>Optional for copying objects in batches (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p16666614131519"><a name="p16666614131519"></a><a name="p16666614131519"></a>Indicates the matching patterns of source objects that are excluded, for example: <strong id="b15405175420523"><a name="b15405175420523"></a><a name="b15405175420523"></a>*.txt</strong>.</p>
<div class="note" id="note145284716208"><a name="note145284716208"></a><a name="note145284716208"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul05604713204"></a><a name="ul05604713204"></a><ul id="ul05604713204"><li>The asterisk (*) represents any group of characters, and the question mark (?) represents any single character. For instance, <strong id="b1955719291487"><a name="b1955719291487"></a><a name="b1955719291487"></a>abc*.txt</strong> indicates any file whose name starts with <strong id="b7558829584"><a name="b7558829584"></a><a name="b7558829584"></a>abc</strong> and ends with <strong id="b185591829186"><a name="b185591829186"></a><a name="b185591829186"></a>.txt</strong>.</li><li>You can use <strong id="b11349193114817"><a name="b11349193114817"></a><a name="b11349193114817"></a>\*</strong> to represent <strong id="b1534953114810"><a name="b1534953114810"></a><a name="b1534953114810"></a>*</strong> and <strong id="b83506311489"><a name="b83506311489"></a><a name="b83506311489"></a>\?</strong> to represent <strong id="b43510311985"><a name="b43510311985"></a><a name="b43510311985"></a>?</strong>.</li><li>If the name of the object to be copied matches the value of this parameter, the object is skipped.</li></ul>
</div></div>
<div class="notice" id="note179117549207"><a name="note179117549207"></a><a name="note179117549207"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><a name="ul877892116516"></a><a name="ul877892116516"></a><ul id="ul877892116516"><li>You are advised to use quotation marks for the matching pattern to prevent special characters from being escaped by the OS and leading to unexpected results. Use single quotation marks for Linux or macOS and quotation marks for Windows.</li><li>The matching pattern applies to the absolute path of an object, including the object name prefix and object name starting from the root directory. For example, if the path of an object in the bucket is <strong id="b4454844474"><a name="b4454844474"></a><a name="b4454844474"></a>obs://bucket/src1/src2/test.txt</strong>, then the absolute path of the object is <strong id="b1445554419719"><a name="b1445554419719"></a><a name="b1445554419719"></a>src1/src2/test.txt</strong>.</li><li>This matching pattern applies only to objects whose names do not end with a slash (/).</li></ul>
</div></div>
</td>
</tr>
<tr id="row207003306314"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p11703123012317"><a name="p11703123012317"></a><a name="p11703123012317"></a>include</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p470523013527"><a name="p470523013527"></a><a name="p470523013527"></a>Optional for copying objects in batches (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p37071730153119"><a name="p37071730153119"></a><a name="p37071730153119"></a>Indicates the matching patterns of source objects that are included, for example: <strong id="b13725131583519"><a name="b13725131583519"></a><a name="b13725131583519"></a>*.jpg</strong>.</p>
<div class="note" id="note195168716220"><a name="note195168716220"></a><a name="note195168716220"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul752013715229"></a><a name="ul752013715229"></a><ul id="ul752013715229"><li>The asterisk (*) represents any group of characters, and the question mark (?) represents any single character.</li><li>You can use <strong id="b85746404813"><a name="b85746404813"></a><a name="b85746404813"></a>\*</strong> to represent <strong id="b757510401984"><a name="b757510401984"></a><a name="b757510401984"></a>*</strong> and <strong id="b1557619401988"><a name="b1557619401988"></a><a name="b1557619401988"></a>\?</strong> to represent <strong id="b17577140386"><a name="b17577140386"></a><a name="b17577140386"></a>?</strong>.</li><li>Only after identifying that the name of the file to be copied does not match the value of <strong id="b31601642282"><a name="b31601642282"></a><a name="b31601642282"></a>exclude</strong>, the system checks whether the file name matches the value of this parameter. If yes, the file is copied. If not, the file is skipped.</li></ul>
</div></div>
<div class="notice" id="note9270217202212"><a name="note9270217202212"></a><a name="note9270217202212"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><a name="ul11495105419578"></a><a name="ul11495105419578"></a><ul id="ul11495105419578"><li>You are advised to use quotation marks for the matching pattern to prevent special characters from being escaped by the OS and leading to unexpected results. Use single quotation marks for Linux or macOS and quotation marks for Windows.</li><li>The matching pattern applies to the absolute path of an object, including the object name prefix and object name starting from the root directory. For example, if the path of an object in the bucket is <strong id="b143935166818"><a name="b143935166818"></a><a name="b143935166818"></a>obs://bucket/src1/src2/test.txt</strong>, then the absolute path of the object is <strong id="b039419161685"><a name="b039419161685"></a><a name="b039419161685"></a>src1/src2/test.txt</strong>.</li><li>This matching pattern applies only to objects whose names do not end with a slash (/).</li></ul>
</div></div>
</td>
</tr>
<tr id="row881016184213"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p496314014481"><a name="p496314014481"></a><a name="p496314014481"></a>timeRange</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p1664510223494"><a name="p1664510223494"></a><a name="p1664510223494"></a>Optional for copying objects in batches (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p16963940134813"><a name="p16963940134813"></a><a name="p16963940134813"></a>Indicates the time range matching pattern when copying objects. Only objects whose last modification time is within the configured time range are copied.</p>
<p id="p1565821819529"><a name="p1565821819529"></a><a name="p1565821819529"></a>This pattern has a lower priority than the object matching patterns (<strong id="b15477142219542"><a name="b15477142219542"></a><a name="b15477142219542"></a>exclude</strong>/<strong id="b104774225546"><a name="b104774225546"></a><a name="b104774225546"></a>include</strong>). That is, the time range matching pattern is executed after the configured object matching patterns.</p>
<div class="note" id="note8766915165919"><a name="note8766915165919"></a><a name="note8766915165919"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul19272430145915"></a><a name="ul19272430145915"></a><ul id="ul19272430145915"><li>The matching time range is represented in <em id="i1525423112818"><a name="i1525423112818"></a><a name="i1525423112818"></a>time1</em><strong id="b6526323172819"><a name="b6526323172819"></a><a name="b6526323172819"></a>-</strong><em id="i3526192318288"><a name="i3526192318288"></a><a name="i3526192318288"></a>time2</em>, where <em id="i1852711231288"><a name="i1852711231288"></a><a name="i1852711231288"></a>time1</em> must be earlier than or the same as <em id="i1952818234281"><a name="i1952818234281"></a><a name="i1952818234281"></a>time2</em>. The time format is <em id="i1752982319285"><a name="i1752982319285"></a><a name="i1752982319285"></a>yyyyMMddHHmmss</em>.</li><li>Automatic formatting is supported. For example, yyyyMMdd is equivalent to yyyyMMdd000000, and yyyyMM is equivalent to yyyyMM01000000.</li><li>If this parameter is set to <strong id="b114382718287"><a name="b114382718287"></a><a name="b114382718287"></a>*-</strong><em id="i17443271283"><a name="i17443271283"></a><a name="i17443271283"></a>time2</em>, all files whose last modification time is earlier than <em id="i245727192819"><a name="i245727192819"></a><a name="i245727192819"></a>time2</em> are matched. If it is set to <em id="i7467273285"><a name="i7467273285"></a><a name="i7467273285"></a>time1</em><strong id="b346427132815"><a name="b346427132815"></a><a name="b346427132815"></a>-*</strong>, all files whose last modification time is later than <em id="i10472279282"><a name="i10472279282"></a><a name="i10472279282"></a>time1</em> are matched.</li></ul>
</div></div>
<div class="notice" id="note1078981817591"><a name="note1078981817591"></a><a name="note1078981817591"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><a name="ul881073612597"></a><a name="ul881073612597"></a><ul id="ul881073612597"><li>Time in the matching pattern is the UTC time.</li><li>This matching pattern applies only to objects whose names do not end with a slash (/).</li></ul>
</div></div>
</td>
</tr>
<tr id="row3161559962"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p10122175516"><a name="p10122175516"></a><a name="p10122175516"></a>mf</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p1562018121550"><a name="p1562018121550"></a><a name="p1562018121550"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p31226718512"><a name="p31226718512"></a><a name="p31226718512"></a>Indicates that the name matching pattern (<strong id="b173026515416"><a name="b173026515416"></a><a name="b173026515416"></a>include</strong> or <strong id="b16219544416"><a name="b16219544416"></a><a name="b16219544416"></a>exclude</strong>) and the time matching pattern (<strong id="b678117911422"><a name="b678117911422"></a><a name="b678117911422"></a>timeRange</strong>) also take effect on objects whose names end with a slash (/).</p>
<p id="p15797105118514"><a name="p15797105118514"></a><a name="p15797105118514"></a></p>
</td>
</tr>
<tr id="row9725193023119"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p1672753013119"><a name="p1672753013119"></a><a name="p1672753013119"></a>o</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p871493095216"><a name="p871493095216"></a><a name="p871493095216"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p137321030143112"><a name="p137321030143112"></a><a name="p137321030143112"></a>Indicates the folder where operation result lists reside. After the command is executed, result lists (possibly including success, failure, and warning files) are generated in the folder. The default value is <strong id="b933316279194"><a name="b933316279194"></a><a name="b933316279194"></a>.obsutil_output</strong>, the subfolder in the home directory of the user who executes obsutil commands.</p>
<div class="note" id="note289083220249"><a name="note289083220249"></a><a name="note289083220249"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul101190347408"></a><a name="ul101190347408"></a><ul id="ul101190347408"><li>The naming rule for result lists is as follows: <strong id="b15699202264514"><a name="b15699202264514"></a><a name="b15699202264514"></a>cp_{succeed  | failed | warning}_report_</strong><em id="i1869919227458"><a name="i1869919227458"></a><a name="i1869919227458"></a>time</em><strong id="b670017222451"><a name="b670017222451"></a><a name="b670017222451"></a>_TaskId.txt</strong></li><li>By default, the maximum size of a single result list is 30 MB and the maximum number of result lists that can be retained is 1024. You can set the maximum size and number by configuring <strong id="b1943164313418"><a name="b1943164313418"></a><a name="b1943164313418"></a>recordMaxLogSize</strong> and <strong id="b174434333412"><a name="b174434333412"></a><a name="b174434333412"></a>recordBackups</strong> in the configuration file.</li></ul>
</div></div>
</td>
</tr>
<tr id="row1248115553510"><td class="cellrowborder" valign="top" width="17.57%" headers="mcps1.1.4.1.1 "><p id="p153951131317"><a name="p153951131317"></a><a name="p153951131317"></a>config</p>
</td>
<td class="cellrowborder" valign="top" width="18.529999999999998%" headers="mcps1.1.4.1.2 "><p id="p12395135316"><a name="p12395135316"></a><a name="p12395135316"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="63.9%" headers="mcps1.1.4.1.3 "><p id="p43952034313"><a name="p43952034313"></a><a name="p43952034313"></a>User-defined configuration file for executing a command. For details about parameters that can be configured, see <a href="parameter-description.md">Parameter Description</a>.</p>
</td>
</tr>
</tbody>
</table>

## Response<a name="section6926520122416"></a>

Refer to  [Response](uploading-an-object.md#section6926520122416)  for uploading an object.

## Running Examples<a name="section1695033152616"></a>

-   Take the Windows OS as an example. Run the  **obsutil cp obs://bucket-test/key obs://bucket-test2**  command to copy a single object.

```
obsutil cp obs://bucket-test/key obs://bucket-test2

Parallel:      3                   Jobs:          3
Threshold:     524288000           PartSize:      5242880
Exclude:                           Include:
VerifyLength:  false               VerifyMd5:     false
CheckpointDir: xxxx

[=====================================================] 100.00% 6/s 0s
Copy successfully, 19B, obs://bucket-test/key --> obs://bucket-test2/key
ext.txt
```

-   Take the Windows OS as an example. Run the  **obsutil cp obs://bucket-test/temp/ obs://bucket-test2 -f -r**  command to copy objects in batches.

```
obsutil cp obs://bucket-test/temp/ obs://bucket-test2 -r -f

Parallel:      3                   Jobs:          3
Threshold:     524288000           PartSize:      5242880
Exclude:                           Include:
VerifyLength:  false               VerifyMd5:     false
CheckpointDir: xxxx
OutputDir: xxxx


[=============================================================] 100.00% 10/s 0s
Succeed count is:   5         Failed count is:    0
Metrics [max cost:298 ms, min cost:192 ms, average cost:238.00 ms, average tps:9.71]
Task id is: 0476929d-9d23-4dc5-b2f8-0a0493f027c5
```

-   For more examples, see  [Copy Examples](copy-examples.md).

