# Listing Buckets<a name="EN-US_TOPIC_0141181363"></a>

## Function<a name="section1479112110815"></a>

You can use this command to obtain the bucket list. In the list, bucket names are displayed in lexicographical order.

## Command Line Structure<a name="section1220752192216"></a>

-   In Windows

    ```
    obsutil ls [-s] [-sc] [-j=1] [-limit=1] [-config=xxx]
    ```

-   In Linux or macOS

    ```
    ./obsutil ls [-s] [-sc] [-j=1] [-limit=1] [-config=xxx]
    ```


## Parameter Description<a name="section6559191102418"></a>

<a name="table10831182114445"></a>
<table><thead align="left"><tr id="row683212154419"><th class="cellrowborder" valign="top" width="18.09%" id="mcps1.1.4.1.1"><p id="p118329219446"><a name="p118329219446"></a><a name="p118329219446"></a>Parameter</p>
</th>
<th class="cellrowborder" valign="top" width="33.18%" id="mcps1.1.4.1.2"><p id="p15137125919108"><a name="p15137125919108"></a><a name="p15137125919108"></a>Optional or Mandatory</p>
</th>
<th class="cellrowborder" valign="top" width="48.730000000000004%" id="mcps1.1.4.1.3"><p id="p12832121184414"><a name="p12832121184414"></a><a name="p12832121184414"></a>Description</p>
</th>
</tr>
</thead>
<tbody><tr id="row167881281811"><td class="cellrowborder" valign="top" width="18.09%" headers="mcps1.1.4.1.1 "><p id="p37889281012"><a name="p37889281012"></a><a name="p37889281012"></a>s</p>
</td>
<td class="cellrowborder" valign="top" width="33.18%" headers="mcps1.1.4.1.2 "><p id="p177888281316"><a name="p177888281316"></a><a name="p177888281316"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="48.730000000000004%" headers="mcps1.1.4.1.3 "><p id="p1378816281111"><a name="p1378816281111"></a><a name="p1378816281111"></a>Displays simplified query result.</p>
<div class="note" id="note1452085114212"><a name="note1452085114212"></a><a name="note1452085114212"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p252035122111"><a name="p252035122111"></a><a name="p252035122111"></a>In the simplified format, the returned result contains only the bucket name.</p>
</div></div>
</td>
</tr>
<tr id="row03718287317"><td class="cellrowborder" valign="top" width="18.09%" headers="mcps1.1.4.1.1 "><p id="p93721828432"><a name="p93721828432"></a><a name="p93721828432"></a>sc</p>
</td>
<td class="cellrowborder" valign="top" width="33.18%" headers="mcps1.1.4.1.2 "><p id="p15372132812313"><a name="p15372132812313"></a><a name="p15372132812313"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="48.730000000000004%" headers="mcps1.1.4.1.3 "><p id="p13372228533"><a name="p13372228533"></a><a name="p13372228533"></a>Queries the storage classes of the buckets when listing buckets.</p>
</td>
</tr>
<tr id="row191337306418"><td class="cellrowborder" valign="top" width="18.09%" headers="mcps1.1.4.1.1 "><p id="p1913323010414"><a name="p1913323010414"></a><a name="p1913323010414"></a>j</p>
</td>
<td class="cellrowborder" valign="top" width="33.18%" headers="mcps1.1.4.1.2 "><p id="p171337309410"><a name="p171337309410"></a><a name="p171337309410"></a>Optional (additional parameter). It must be used together with <strong id="b770944718612"><a name="b770944718612"></a><a name="b770944718612"></a>sc</strong>.</p>
</td>
<td class="cellrowborder" valign="top" width="48.730000000000004%" headers="mcps1.1.4.1.3 "><p id="p1513363013419"><a name="p1513363013419"></a><a name="p1513363013419"></a>Indicates the maximum number of concurrent tasks for querying the bucket storage class. The default value is the value of <strong id="b129310433513"><a name="b129310433513"></a><a name="b129310433513"></a>defaultJobs</strong> in the configuration file.</p>
<div class="note" id="note891964620819"><a name="note891964620819"></a><a name="note891964620819"></a><span class="notetitle"> NOTE: </span><div class="notebody"><p id="p1091964618820"><a name="p1091964618820"></a><a name="p1091964618820"></a>The value is ensured to be greater than or equal to 1.</p>
</div></div>
</td>
</tr>
<tr id="row918019240510"><td class="cellrowborder" valign="top" width="18.09%" headers="mcps1.1.4.1.1 "><p id="p1918111241654"><a name="p1918111241654"></a><a name="p1918111241654"></a>limit</p>
</td>
<td class="cellrowborder" valign="top" width="33.18%" headers="mcps1.1.4.1.2 "><p id="p218113247513"><a name="p218113247513"></a><a name="p218113247513"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="48.730000000000004%" headers="mcps1.1.4.1.3 "><p id="p310214615616"><a name="p310214615616"></a><a name="p310214615616"></a>Maximum number of buckets that can be queried. If the value is less than 0, all buckets are listed. If it is left blank, a maximum of 1000 buckets can be listed by default.</p>
</td>
</tr>
<tr id="row617302812339"><td class="cellrowborder" valign="top" width="18.09%" headers="mcps1.1.4.1.1 "><p id="p153951131317"><a name="p153951131317"></a><a name="p153951131317"></a>config</p>
</td>
<td class="cellrowborder" valign="top" width="33.18%" headers="mcps1.1.4.1.2 "><p id="p12395135316"><a name="p12395135316"></a><a name="p12395135316"></a>Optional (additional parameter)</p>
</td>
<td class="cellrowborder" valign="top" width="48.730000000000004%" headers="mcps1.1.4.1.3 "><p id="p43952034313"><a name="p43952034313"></a><a name="p43952034313"></a>User-defined configuration file for executing a command. For details about parameters that can be configured, see <a href="parameter-description.md">Parameter Description</a>.</p>
</td>
</tr>
</tbody>
</table>

## Running Example<a name="section15899161919244"></a>

-   Take the Windows OS as an example. Run the  **obsutil ls -limit=5**  command to obtain the bucket list.

```
obsutil ls -limit=5

Bucket                        CreationDate                  Location         BucketType   
obs://bucket001               2018-09-03T01:53:02Z          example          OBJECT   

obs://bucket002               2018-11-01T01:40:01Z          example          OBJECT      

obs://bucket003               2018-10-25T11:45:45Z          example          OBJECT       

obs://bucket004               2018-10-26T02:33:09Z          example          OBJECT      

obs://bucket005               2018-10-26T02:34:50Z          example          OBJECT      
         
Bucket number is: 5
```

>![](public_sys-resources/icon-note.gif) **NOTE:**   
>In the bucket listing result, the  **BucketType**  field indicates the bucket type:  **OBJECT**  indicates the bucket for object storage;  **POSIX**  indicates the parallel file system.  

