# Resuming a Failed Copy Task<a name="EN-US_TOPIC_0145591177"></a>

## Function<a name="section1479112110815"></a>

You can use this command to resume a failed copy task based on the task ID.

## Command Line Structure<a name="section49408320267"></a>

-   In Windows

    ```
    obsutil cp -recover=xxx [-dryRun] [-f] [-u] [-crr] [-vlength] [-vmd5] [-j=1] [-p=1] [-threshold=52428800] [-acl=xxx] [-sc=xxx] [-meta=aaa:bbb#ccc:ddd] [-ps=auto] [-include=*.xxx] [-exclude=*.xxx] [-timeRange=time1-time2] [-mf] [-o=xxx] [-cpd=xxx] [-config=xxx]
    ```

-   In Linux or macOS

    ```
    ./obsutil cp -recover=xxx [-dryRun] [-f] [-u] [-crr] [-vlength] [-vmd5] [-j=1] [-p=1] [-threshold=52428800] [-acl=xxx] [-sc=xxx] [-meta=aaa:bbb#ccc:ddd] [-ps=auto] [-include=*.xxx] [-exclude=*.xxx] [-timeRange=time1-time2] [-mf] [-o=xxx] [-cpd=xxx] [-config=xxx]
    ```


## Parameter Description<a name="section3940203142610"></a>

<a name="table125609307312"></a>
<table><thead align="left"><tr id="row456603043117"><th class="cellrowborder" valign="top" width="18%" id="mcps1.1.4.1.1"><p id="p85680303316"><a name="p85680303316"></a><a name="p85680303316"></a>Parameter</p>
</th>
<th class="cellrowborder" valign="top" width="27%" id="mcps1.1.4.1.2"><p id="p14572130153112"><a name="p14572130153112"></a><a name="p14572130153112"></a>Optional or Mandatory</p>
</th>
<th class="cellrowborder" valign="top" width="55.00000000000001%" id="mcps1.1.4.1.3"><p id="p15751930133118"><a name="p15751930133118"></a><a name="p15751930133118"></a>Description</p>
</th>
</tr>
</thead>
<tbody><tr id="row10984414202"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p28307351476"><a name="p28307351476"></a><a name="p28307351476"></a>recover</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p198304351672"><a name="p198304351672"></a><a name="p198304351672"></a>Mandatory (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p188308358715"><a name="p188308358715"></a><a name="p188308358715"></a>ID of the copy task to be resumed.</p>
<div class="note" id="note10332165610814"><a name="note10332165610814"></a><a name="note10332165610814"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul625425894"></a><a name="ul625425894"></a><ul id="ul625425894"><li>You can obtain the task ID after a copy task is complete, or query it based on the file name of the operation result list, which is the 36 characters excluding the suffix <strong id="b524681817582"><a name="b524681817582"></a><a name="b524681817582"></a>.txt</strong> in the file name.</li><li>You can locate the copy task to be resumed in the directory where the result lists reside. For details about the directory of the result lists, see additional parameter <strong id="b1613810603917"><a name="b1613810603917"></a><a name="b1613810603917"></a>o</strong>.</li></ul>
</div></div>
</td>
</tr>
<tr id="row1335717441023"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p156631810914"><a name="p156631810914"></a><a name="p156631810914"></a>dryRun</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p10663610617"><a name="p10663610617"></a><a name="p10663610617"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p76630101412"><a name="p76630101412"></a><a name="p76630101412"></a>Conducts a dry run.</p>
</td>
</tr>
<tr id="row12398958133217"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p19960122065918"><a name="p19960122065918"></a><a name="p19960122065918"></a>crr</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p696012075919"><a name="p696012075919"></a><a name="p696012075919"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p1296062012595"><a name="p1296062012595"></a><a name="p1296062012595"></a>Enables the client-side cross-region replication function. In this mode, data is directly copied to the destination bucket from the source bucket through data stream. The buckets can by any two OBS buckets.</p>
<div class="note" id="note4424164104813"><a name="note4424164104813"></a><a name="note4424164104813"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul743879115112"></a><a name="ul743879115112"></a><ul id="ul743879115112"><li>If this parameter is configured, ensure that the configuration of client-side cross-region replication is updated in the configuration file. For details, see <a href="updating-a-configuration-file.md">Updating a Configuration File</a>.</li><li>The configurations of the source bucket and destination bucket are respectively <strong id="b1028317394322"><a name="b1028317394322"></a><a name="b1028317394322"></a>akCrr/skCrr/tokenCrr/endpointCrr</strong> and <strong id="b19283203973211"><a name="b19283203973211"></a><a name="b19283203973211"></a>ak/sk/token/endpoint</strong> in the configuration file.</li></ul>
</div></div>
<div class="notice" id="note19922116154811"><a name="note19922116154811"></a><a name="note19922116154811"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><p id="p4922106154811"><a name="p4922106154811"></a><a name="p4922106154811"></a>After this function is enabled, both upload and download bandwidth are occupied.</p>
</div></div>
</td>
</tr>
<tr id="row69712557325"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p1605745105311"><a name="p1605745105311"></a><a name="p1605745105311"></a>vlength</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p12605114512534"><a name="p12605114512534"></a><a name="p12605114512534"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p1660724595316"><a name="p1660724595316"></a><a name="p1660724595316"></a>Verifies whether the object size in the destination bucket is the same as that in the source bucket after the copy task completes.</p>
<div class="note" id="note57671113172915"><a name="note57671113172915"></a><a name="note57671113172915"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p1768313182910"><a name="p1768313182910"></a><a name="p1768313182910"></a>This parameter must be used together with <strong id="b122200479322"><a name="b122200479322"></a><a name="b122200479322"></a>crr</strong>.</p>
</div></div>
</td>
</tr>
<tr id="row1360005310322"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p136091945105315"><a name="p136091945105315"></a><a name="p136091945105315"></a>vmd5</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p461011452539"><a name="p461011452539"></a><a name="p461011452539"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p19611745185310"><a name="p19611745185310"></a><a name="p19611745185310"></a>Verifies whether the MD5 value of the destination bucket is the same as that of the source bucket after the copy task completes.</p>
<div class="note" id="note13542169102916"><a name="note13542169102916"></a><a name="note13542169102916"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul114913379291"></a><a name="ul114913379291"></a><ul id="ul114913379291"><li>This parameter must be used together with <strong id="b18530115119329"><a name="b18530115119329"></a><a name="b18530115119329"></a>crr</strong>.</li><li>Objects in the source bucket must contain metadata <strong id="b39201852143215"><a name="b39201852143215"></a><a name="b39201852143215"></a>x-obs-md5chksum</strong>. Otherwise, MD5 verification will be skipped.<p id="p1765649132314"><a name="p1765649132314"></a><a name="p1765649132314"></a>After the MD5 value verification is successful, the parameter value is set to the destination object metadata <strong id="b570344710718"><a name="b570344710718"></a><a name="b570344710718"></a>x-obs-md5chksum</strong>, which is used for later MD5 verification during download or copy.</p>
</li></ul>
</div></div>
</td>
</tr>
<tr id="row17612113018317"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p1461418302318"><a name="p1461418302318"></a><a name="p1461418302318"></a>u</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p1615183013110"><a name="p1615183013110"></a><a name="p1615183013110"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p361753019315"><a name="p361753019315"></a><a name="p361753019315"></a>Indicates incremental copy. If this parameter is set, each object can be copied only when it does not exist in the destination bucket, its size is different from the namesake one in the destination bucket, or it has the latest modification time.</p>
</td>
</tr>
<tr id="row3632193093117"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p26341030163117"><a name="p26341030163117"></a><a name="p26341030163117"></a>p</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p9636183093120"><a name="p9636183093120"></a><a name="p9636183093120"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p7638163013112"><a name="p7638163013112"></a><a name="p7638163013112"></a>Indicates the maximum number of concurrent multipart copy tasks when copying an object. The default value is the value of <strong id="b1487142617214"><a name="b1487142617214"></a><a name="b1487142617214"></a>defaultParallels</strong> in the configuration file.</p>
</td>
</tr>
<tr id="row363913014313"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p116401030183113"><a name="p116401030183113"></a><a name="p116401030183113"></a>threshold</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p176431304319"><a name="p176431304319"></a><a name="p176431304319"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p16441830113114"><a name="p16441830113114"></a><a name="p16441830113114"></a>Indicates the threshold for enabling multipart copy, in bytes. The default value is the value of <strong id="b289264273012"><a name="b289264273012"></a><a name="b289264273012"></a>defaultBigfileThreshold</strong> in the configuration file.</p>
<div class="note" id="note1382352674913"><a name="note1382352674913"></a><a name="note1382352674913"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul182059695115"></a><a name="ul182059695115"></a><ul id="ul182059695115"><li>If the size of the object to be copied is smaller than the threshold, copy the object directly. If not, a multipart copy is required.</li><li>If you copy an object directly, no part record is generated, and resumable transmission is not supported.</li><li>This value can contain a capacity unit. For example, <strong id="b1681629194512"><a name="b1681629194512"></a><a name="b1681629194512"></a>1 MB</strong> indicates 1048576 bytes.</li></ul>
</div></div>
</td>
</tr>
<tr id="row264453012315"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p964614304315"><a name="p964614304315"></a><a name="p964614304315"></a>acl</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p17648193011318"><a name="p17648193011318"></a><a name="p17648193011318"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p16650163083117"><a name="p16650163083117"></a><a name="p16650163083117"></a>Access control policies for destination objects that can be specified when copying objects. Possible values are:</p>
<a name="ul0651183053115"></a><a name="ul0651183053115"></a><ul id="ul0651183053115"><li>private</li><li>public-read</li><li>public-read-write</li></ul>
<div class="note" id="note1790113183525"><a name="note1790113183525"></a><a name="note1790113183525"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p92982882916"><a name="p92982882916"></a><a name="p92982882916"></a>The preceding three values indicate private read and write, public read, and public read and write.</p>
</div></div>
</td>
</tr>
<tr id="row62441623105116"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p13466172515511"><a name="p13466172515511"></a><a name="p13466172515511"></a>sc</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p134681525175117"><a name="p134681525175117"></a><a name="p134681525175117"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p04731025135114"><a name="p04731025135114"></a><a name="p04731025135114"></a>Storage classes of the destination objects that can be specified when copying objects. Possible values are:</p>
<a name="ul175651814214"></a><a name="ul175651814214"></a><ul id="ul175651814214"><li><strong id="b14949324202416"><a name="b14949324202416"></a><a name="b14949324202416"></a>standard</strong>: OBS Standard, which features low access latency and high throughput, and is applicable to storing frequently accessed data (multiple accesses per month averagely) or data that is smaller than 1 MB</li><li><strong id="b13234131410817"><a name="b13234131410817"></a><a name="b13234131410817"></a>warm</strong>: OBS Warm. It is applicable to storing semi-frequently accessed (less than 12 times a year averagely) data that requires quick response.</li><li><strong id="b168399185818"><a name="b168399185818"></a><a name="b168399185818"></a>cold</strong>: OBS Cold. It is secure, durable, and inexpensive, and applicable to archiving rarely-accessed (once a year averagely) data.</li></ul>
</td>
</tr>
<tr id="row8659203023112"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p1366223093111"><a name="p1366223093111"></a><a name="p1366223093111"></a>meta</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p26631230193118"><a name="p26631230193118"></a><a name="p26631230193118"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p17941525135111"><a name="p17941525135111"></a><a name="p17941525135111"></a>Metadata of destination objects that can be specified when copying objects. The format is <em id="i82814581343"><a name="i82814581343"></a><a name="i82814581343"></a>key1</em><strong id="b2028235814414"><a name="b2028235814414"></a><a name="b2028235814414"></a>:</strong><em id="i102823581549"><a name="i102823581549"></a><a name="i102823581549"></a>value1</em><strong id="b228310589414"><a name="b228310589414"></a><a name="b228310589414"></a>#</strong><em id="i1728413581241"><a name="i1728413581241"></a><a name="i1728413581241"></a>key2</em><strong id="b1528465812419"><a name="b1528465812419"></a><a name="b1528465812419"></a>:</strong><em id="i16285658144"><a name="i16285658144"></a><a name="i16285658144"></a>value2</em><strong id="b528613581417"><a name="b528613581417"></a><a name="b528613581417"></a>#</strong><em id="i1128618583418"><a name="i1128618583418"></a><a name="i1128618583418"></a>key3</em><strong id="b132870589414"><a name="b132870589414"></a><a name="b132870589414"></a>:</strong><em id="i82878581941"><a name="i82878581941"></a><a name="i82878581941"></a>value3</em>.</p>
<div class="note" id="note1606229165119"><a name="note1606229165119"></a><a name="note1606229165119"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p686342013559"><a name="p686342013559"></a><a name="p686342013559"></a>The preceding value indicates that the destination objects in the bucket contain three groups of customized metadata after objects are copied: <strong id="b265591172211"><a name="b265591172211"></a><a name="b265591172211"></a>key1:value1</strong>, <strong id="b156576116225"><a name="b156576116225"></a><a name="b156576116225"></a>key2:value2</strong>, and <strong id="b1665813112223"><a name="b1665813112223"></a><a name="b1665813112223"></a>key3:value3</strong>.</p>
</div></div>
</td>
</tr>
<tr id="row466716303317"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p1366917307314"><a name="p1366917307314"></a><a name="p1366917307314"></a>ps</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p267203023115"><a name="p267203023115"></a><a name="p267203023115"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p1167313023117"><a name="p1167313023117"></a><a name="p1167313023117"></a>Indicates the size of each part in a multipart copy task, in bytes. The value ranges from 100 KB to 5 GB. The default value is the value of <strong id="b1782672311313"><a name="b1782672311313"></a><a name="b1782672311313"></a>defaultPartSize</strong> in the configuration file.</p>
<div class="note" id="note19227113864316"><a name="note19227113864316"></a><a name="note19227113864316"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul11281315134615"></a><a name="ul11281315134615"></a><ul id="ul11281315134615"><li>This value can contain a capacity unit. For example, <strong id="b9538193835216"><a name="b9538193835216"></a><a name="b9538193835216"></a>1 MB</strong> indicates 1048576 bytes.</li><li>The parameter can be set to <strong id="b144601040105219"><a name="b144601040105219"></a><a name="b144601040105219"></a>auto</strong>. In this case, obsutil automatically sets the part size for each multipart task based on the source object size.</li></ul>
</div></div>
</td>
</tr>
<tr id="row1367463093119"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p1967612305318"><a name="p1967612305318"></a><a name="p1967612305318"></a>cpd</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p1367810308316"><a name="p1367810308316"></a><a name="p1367810308316"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p46792307316"><a name="p46792307316"></a><a name="p46792307316"></a>Indicates the folder where the part records reside. The default value is <strong id="b652994419198"><a name="b652994419198"></a><a name="b652994419198"></a>.obsutil_checkpoint</strong>, the subfolder in the home directory of the user who executes obsutil commands.</p>
<div class="note" id="note22601842192718"><a name="note22601842192718"></a><a name="note22601842192718"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p118866577266"><a name="p118866577266"></a><a name="p118866577266"></a>A part record is generated during a multipart copy and saved to the <strong id="b2029618267491"><a name="b2029618267491"></a><a name="b2029618267491"></a>copy</strong> subfolder. After the copy succeeds, its part record is deleted automatically. If the copy fails or is suspended, the system attempts to resume the task according to its part record when you perform the copy the next time.</p>
</div></div>
</td>
</tr>
<tr id="row1568763019310"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p2688183018315"><a name="p2688183018315"></a><a name="p2688183018315"></a>f</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p1748156161912"><a name="p1748156161912"></a><a name="p1748156161912"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p1269315302318"><a name="p1269315302318"></a><a name="p1269315302318"></a>Runs in force mode.</p>
</td>
</tr>
<tr id="row20694163013111"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p269613053110"><a name="p269613053110"></a><a name="p269613053110"></a>j</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p177527681919"><a name="p177527681919"></a><a name="p177527681919"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p4699130153112"><a name="p4699130153112"></a><a name="p4699130153112"></a>Indicates the maximum number of concurrent tasks for copying objects in batches. The default value is the value of <strong id="b9810812011"><a name="b9810812011"></a><a name="b9810812011"></a>defaultJobs</strong> in the configuration file.</p>
<div class="note" id="note891964620819"><a name="note891964620819"></a><a name="note891964620819"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p1091964618820"><a name="p1091964618820"></a><a name="p1091964618820"></a>The value is ensured to be greater than or equal to 1.</p>
</div></div>
</td>
</tr>
<tr id="row8334161131517"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p7660121421518"><a name="p7660121421518"></a><a name="p7660121421518"></a>exclude</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p207547618193"><a name="p207547618193"></a><a name="p207547618193"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p16666614131519"><a name="p16666614131519"></a><a name="p16666614131519"></a>Indicates the matching patterns of source objects that are excluded, for example: <strong id="b15405175420523"><a name="b15405175420523"></a><a name="b15405175420523"></a>*.txt</strong>.</p>
<div class="note" id="note145284716208"><a name="note145284716208"></a><a name="note145284716208"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul05604713204"></a><a name="ul05604713204"></a><ul id="ul05604713204"><li>The asterisk (*) represents any group of characters, and the question mark (?) represents any single character. For instance, <strong id="b2768131118911"><a name="b2768131118911"></a><a name="b2768131118911"></a>abc*.txt</strong> indicates any file whose name starts with <strong id="b167705116919"><a name="b167705116919"></a><a name="b167705116919"></a>abc</strong> and ends with <strong id="b277213112098"><a name="b277213112098"></a><a name="b277213112098"></a>.txt</strong>.</li><li>You can use <strong id="b1319617136917"><a name="b1319617136917"></a><a name="b1319617136917"></a>\*</strong> to represent <strong id="b1419715131995"><a name="b1419715131995"></a><a name="b1419715131995"></a>*</strong> and <strong id="b31983131795"><a name="b31983131795"></a><a name="b31983131795"></a>\?</strong> to represent <strong id="b12001513198"><a name="b12001513198"></a><a name="b12001513198"></a>?</strong>.</li><li>If the name of the object to be copied matches the value of this parameter, the object is skipped.</li></ul>
</div></div>
<div class="notice" id="note179117549207"><a name="note179117549207"></a><a name="note179117549207"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><a name="ul877892116516"></a><a name="ul877892116516"></a><ul id="ul877892116516"><li>You are advised to use quotation marks for the matching pattern to prevent special characters from being escaped by the OS and leading to unexpected results. Use single quotation marks for Linux or macOS and quotation marks for Windows.</li><li>The matching pattern applies to the absolute path of an object, including the object name prefix and object name starting from the root directory. For example, if the path of an object in the bucket is <strong id="b114981623151511"><a name="b114981623151511"></a><a name="b114981623151511"></a>obs://bucket/src1/src2/test.txt</strong>, then the absolute path of the object is <strong id="b1449912316155"><a name="b1449912316155"></a><a name="b1449912316155"></a>src1/src2/test.txt</strong>.</li><li>This matching pattern applies only to objects whose names do not end with a slash (/).</li></ul>
</div></div>
</td>
</tr>
<tr id="row207003306314"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p11703123012317"><a name="p11703123012317"></a><a name="p11703123012317"></a>include</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p1275712617196"><a name="p1275712617196"></a><a name="p1275712617196"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p37071730153119"><a name="p37071730153119"></a><a name="p37071730153119"></a>Indicates the matching patterns of source objects that are included, for example: <strong id="b13725131583519"><a name="b13725131583519"></a><a name="b13725131583519"></a>*.jpg</strong>.</p>
<div class="note" id="note195168716220"><a name="note195168716220"></a><a name="note195168716220"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul752013715229"></a><a name="ul752013715229"></a><ul id="ul752013715229"><li>The asterisk (*) represents any group of characters, and the question mark (?) represents any single character.</li><li>You can use <strong id="b675322113915"><a name="b675322113915"></a><a name="b675322113915"></a>\*</strong> to represent <strong id="b117544215918"><a name="b117544215918"></a><a name="b117544215918"></a>*</strong> and <strong id="b187551211199"><a name="b187551211199"></a><a name="b187551211199"></a>\?</strong> to represent <strong id="b075672112917"><a name="b075672112917"></a><a name="b075672112917"></a>?</strong>.</li><li>Only after identifying that the name of the file to be copied does not match the value of <strong id="b135951230920"><a name="b135951230920"></a><a name="b135951230920"></a>exclude</strong>, the system checks whether the file name matches the value of this parameter. If yes, the file is copied. If not, the file is skipped.</li></ul>
</div></div>
<div class="notice" id="note9270217202212"><a name="note9270217202212"></a><a name="note9270217202212"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><a name="ul1514124145810"></a><a name="ul1514124145810"></a><ul id="ul1514124145810"><li>You are advised to use quotation marks for the matching pattern to prevent special characters from being escaped by the OS and leading to unexpected results. Use single quotation marks for Linux or macOS and quotation marks for Windows.</li><li>The matching pattern applies to the absolute path of an object, including the object name prefix and object name starting from the root directory. For example, if the path of an object in the bucket is <strong id="b1674042471518"><a name="b1674042471518"></a><a name="b1674042471518"></a>obs://bucket/src1/src2/test.txt</strong>, then the absolute path of the object is <strong id="b12741112415152"><a name="b12741112415152"></a><a name="b12741112415152"></a>src1/src2/test.txt</strong>.</li><li>This matching pattern applies only to objects whose names do not end with a slash (/).</li></ul>
</div></div>
</td>
</tr>
<tr id="row450632218429"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p496314014481"><a name="p496314014481"></a><a name="p496314014481"></a>timeRange</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p1664510223494"><a name="p1664510223494"></a><a name="p1664510223494"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p16963940134813"><a name="p16963940134813"></a><a name="p16963940134813"></a>Indicates the time range matching pattern when copying objects. Only objects whose last modification time is within the configured time range are copied.</p>
<p id="p169680339524"><a name="p169680339524"></a><a name="p169680339524"></a>This pattern has a lower priority than the object matching patterns (<strong id="b14528969560"><a name="b14528969560"></a><a name="b14528969560"></a>exclude</strong>/<strong id="b6529116105611"><a name="b6529116105611"></a><a name="b6529116105611"></a>include</strong>). That is, the time range matching pattern is executed after the configured object matching patterns.</p>
<div class="note" id="note8766915165919"><a name="note8766915165919"></a><a name="note8766915165919"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul19272430145915"></a><a name="ul19272430145915"></a><ul id="ul19272430145915"><li>The matching time range is represented in <em id="i12207102781515"><a name="i12207102781515"></a><a name="i12207102781515"></a>time1</em><strong id="b12208162761517"><a name="b12208162761517"></a><a name="b12208162761517"></a>-</strong><em id="i152091827111518"><a name="i152091827111518"></a><a name="i152091827111518"></a>time2</em>, where <em id="i11209827131514"><a name="i11209827131514"></a><a name="i11209827131514"></a>time1</em> must be earlier than or the same as <em id="i321015277151"><a name="i321015277151"></a><a name="i321015277151"></a>time2</em>. The time format is <em id="i1821017278155"><a name="i1821017278155"></a><a name="i1821017278155"></a>yyyyMMddHHmmss</em>.</li><li>Automatic formatting is supported. For example, yyyyMMdd is equivalent to yyyyMMdd000000, and yyyyMM is equivalent to yyyyMM01000000.</li><li>If this parameter is set to <strong id="b092117442158"><a name="b092117442158"></a><a name="b092117442158"></a>*-</strong><em id="i9922644161514"><a name="i9922644161514"></a><a name="i9922644161514"></a>time2</em>, all files whose last modification time is earlier than <em id="i6922644151517"><a name="i6922644151517"></a><a name="i6922644151517"></a>time2</em> are matched. If it is set to <em id="i1992354401516"><a name="i1992354401516"></a><a name="i1992354401516"></a>time1</em><strong id="b1392364471518"><a name="b1392364471518"></a><a name="b1392364471518"></a>-*</strong>, all files whose last modification time is later than <em id="i69244440157"><a name="i69244440157"></a><a name="i69244440157"></a>time1</em> are matched.</li></ul>
</div></div>
<div class="notice" id="note1078981817591"><a name="note1078981817591"></a><a name="note1078981817591"></a><span class="noticetitle"> NOTICE: </span><div class="noticebody"><a name="ul881073612597"></a><a name="ul881073612597"></a><ul id="ul881073612597"><li>Time in the matching pattern is the UTC time.</li><li>This matching pattern applies only to objects whose names do not end with a slash (/).</li></ul>
</div></div>
</td>
</tr>
<tr id="row2100330397"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p10122175516"><a name="p10122175516"></a><a name="p10122175516"></a>mf</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p1562018121550"><a name="p1562018121550"></a><a name="p1562018121550"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p31226718512"><a name="p31226718512"></a><a name="p31226718512"></a>Indicates that the name matching pattern (<strong id="b17339350857"><a name="b17339350857"></a><a name="b17339350857"></a>include</strong> or <strong id="b1834017508517"><a name="b1834017508517"></a><a name="b1834017508517"></a>exclude</strong>) and the time matching pattern (<strong id="b5341145012515"><a name="b5341145012515"></a><a name="b5341145012515"></a>timeRange</strong>) also take effect on objects whose names end with a slash (/).</p>
<p id="p15797105118514"><a name="p15797105118514"></a><a name="p15797105118514"></a></p>
</td>
</tr>
<tr id="row9725193023119"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p1672753013119"><a name="p1672753013119"></a><a name="p1672753013119"></a>o</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p17760136151914"><a name="p17760136151914"></a><a name="p17760136151914"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p137321030143112"><a name="p137321030143112"></a><a name="p137321030143112"></a>Indicates the folder where operation result lists reside. After the command is executed, result lists (possibly including success, failure, and warning files) are generated in the folder. The default value is <strong id="b1398819332010"><a name="b1398819332010"></a><a name="b1398819332010"></a>.obsutil_output</strong>, the subfolder in the home directory of the user who executes obsutil commands.</p>
<div class="note" id="note289083220249"><a name="note289083220249"></a><a name="note289083220249"></a><span class="notetitle"> NOTE: </span><div class="notebody"><a name="ul101190347408"></a><a name="ul101190347408"></a><ul id="ul101190347408"><li>The naming rule for result lists is as follows: <strong id="b1144671115334"><a name="b1144671115334"></a><a name="b1144671115334"></a>cp_{succeed  | failed | warning}_report_</strong><em id="i9447131110335"><a name="i9447131110335"></a><a name="i9447131110335"></a>time</em><strong id="b04471411123313"><a name="b04471411123313"></a><a name="b04471411123313"></a>_TaskId.txt</strong><p id="li412103664016p0"><a name="li412103664016p0"></a><a name="li412103664016p0"></a>By default, the maximum size of a single result list is 30 MB and the maximum number of result lists that can be retained is 1024. You can set the maximum size and number by configuring <strong id="b1947712185392"><a name="b1947712185392"></a><a name="b1947712185392"></a>recordMaxLogSize</strong> and <strong id="b74781818193914"><a name="b74781818193914"></a><a name="b74781818193914"></a>recordBackups</strong> in the configuration file.</p>
</li></ul>
</div></div>
</td>
</tr>
<tr id="row923556133614"><td class="cellrowborder" valign="top" width="18%" headers="mcps1.1.4.1.1 "><p id="p153951131317"><a name="p153951131317"></a><a name="p153951131317"></a>config</p>
</td>
<td class="cellrowborder" valign="top" width="27%" headers="mcps1.1.4.1.2 "><p id="p12395135316"><a name="p12395135316"></a><a name="p12395135316"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="55.00000000000001%" headers="mcps1.1.4.1.3 "><p id="p43952034313"><a name="p43952034313"></a><a name="p43952034313"></a>User-defined configuration file for executing a command. For details about parameters that can be configured, see <a href="parameter-description.md">Parameter Description</a>.</p>
</td>
</tr>
</tbody>
</table>

## Response<a name="section6926520122416"></a>

Refer to  [Response](uploading-an-object.md#section6926520122416)  for uploading an object.

## Running Example<a name="section1695033152616"></a>

-   Take the Windows OS as an example. Run the  **obsutil cp -recover=0476929d-9d23-4dc5-b2f8-0a0493f027c5 -f**  command to copy objects in batches.

```
obsutil cp -recover=0476929d-9d23-4dc5-b2f8-0a0493f027c5 -f

Parallel:      3                   Jobs:          3
Threshold:     524288000           PartSize:      5242880
Exclude:                           Include:
VerifyLength:  false               VerifyMd5:     false
CheckpointDir: xxxx
OutputDir: xxxx


[=============================================================] 100.00% 10/s 0s
Succeed count is:   1         Failed count is:    0
Metrics [max cost:298 ms, min cost:192 ms, average cost:238.00 ms, average tps:9.71]
Task id is: f4c4f2b6-6e54-4dff-96b8-52e8c8c9a4b0
```

