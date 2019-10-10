# Viewing Command Help Information<a name="EN-US_TOPIC_0141181371"></a>

## Function<a name="section1479112110815"></a>

You can use this command to view the commands supported by obsutil or view the help information of a specific command.

## Command Line Structure<a name="section1220752192216"></a>

-   In Windows

    ```
    obsutil help [command]
    ```

-   In Linux or macOS

    ```
    ./obsutil help [command]
    ```


## Parameter Description<a name="section6559191102418"></a>

<a name="table10831182114445"></a>
<table><thead align="left"><tr id="row683212154419"><th class="cellrowborder" valign="top" width="17%" id="mcps1.1.4.1.1"><p id="p118329219446"><a name="p118329219446"></a><a name="p118329219446"></a>Parameter</p>
</th>
<th class="cellrowborder" valign="top" width="26%" id="mcps1.1.4.1.2"><p id="p15137125919108"><a name="p15137125919108"></a><a name="p15137125919108"></a>Optional or Mandatory</p>
</th>
<th class="cellrowborder" valign="top" width="56.99999999999999%" id="mcps1.1.4.1.3"><p id="p12832121184414"><a name="p12832121184414"></a><a name="p12832121184414"></a>Description</p>
</th>
</tr>
</thead>
<tbody><tr id="row108328217449"><td class="cellrowborder" valign="top" width="17%" headers="mcps1.1.4.1.1 "><p id="p64495172515"><a name="p64495172515"></a><a name="p64495172515"></a>command</p>
</td>
<td class="cellrowborder" valign="top" width="26%" headers="mcps1.1.4.1.2 "><p id="p154316502519"><a name="p154316502519"></a><a name="p154316502519"></a>Optional</p>
</td>
<td class="cellrowborder" valign="top" width="56.99999999999999%" headers="mcps1.1.4.1.3 "><p id="p17425512259"><a name="p17425512259"></a><a name="p17425512259"></a>Currently, the help documents of the following commands are available:</p>
<a name="ul921413323261"></a><a name="ul921413323261"></a><ul id="ul921413323261"><li>For <strong id="b4463635255"><a name="b4463635255"></a><a name="b4463635255"></a>abort</strong>, see <a href="deleting-a-multipart-upload-task.md">Deleting a Multipart Upload Task</a>.</li><li>For <strong id="b15591439153"><a name="b15591439153"></a><a name="b15591439153"></a>chattri</strong>, see <a href="setting-object-properties.md">Setting Object Properties</a>.</li><li>For <strong id="b186289502518"><a name="b186289502518"></a><a name="b186289502518"></a>cp</strong>, see <a href="uploading-an-object.md">Uploading an Object</a>, <a href="copying-an-object.md">Copying an Object</a>, and <a href="downloading-an-object.md">Downloading an Object</a>.</li><li>For <strong id="b20312126163"><a name="b20312126163"></a><a name="b20312126163"></a>ls</strong>, see <a href="listing-buckets.md">Listing Buckets</a>, <a href="listing-objects.md">Listing Objects</a>, and <a href="listing-multipart-upload-tasks.md">Listing Multipart Upload Tasks</a>.</li><li>For <strong id="b1772048461"><a name="b1772048461"></a><a name="b1772048461"></a>mb</strong>, see <a href="creating-a-bucket.md">Creating a Bucket</a>.</li><li>For <strong id="b13818151871"><a name="b13818151871"></a><a name="b13818151871"></a>mkdir</strong>, see <a href="creating-a-folder.md">Creating a Folder</a>.</li><li>For <strong id="b51386438711"><a name="b51386438711"></a><a name="b51386438711"></a>mv</strong>, see <a href="moving-an-object.md">Moving an Object</a>.</li><li>For <strong id="b1711915413813"><a name="b1711915413813"></a><a name="b1711915413813"></a>restore</strong>, see <a href="restoring-objects-from-obs-cold.md">Restoring Objects from OBS Cold</a>.</li><li>For <strong id="b42066221885"><a name="b42066221885"></a><a name="b42066221885"></a>rm</strong>, see <a href="deleting-a-bucket.md">Deleting a Bucket</a> and <a href="deleting-an-object.md">Deleting an Object</a>.</li><li>For <strong id="b198165495817"><a name="b198165495817"></a><a name="b198165495817"></a>sign</strong>, see <a href="generating-the-download-link-of-an-object.md">Generating the Download Link of an Object</a>.</li><li>For <strong id="b20382112092"><a name="b20382112092"></a><a name="b20382112092"></a>stat</strong>, see <a href="querying-bucket-properties.md">Querying Bucket Properties</a> and <a href="querying-object-properties.md">Querying Object Properties</a>.</li><li>For <strong id="b139191232493"><a name="b139191232493"></a><a name="b139191232493"></a>sync</strong>, see <a href="synchronously-uploading-incremental-objects.md">Synchronously Uploading Incremental Objects</a>, <a href="synchronously-copying-incremental-objects.md">Synchronously Copying Incremental Objects</a>, and <a href="synchronously-downloading-incremental-objects.md">Synchronously Downloading Incremental Objects</a>.</li><li>For <strong id="b584712761019"><a name="b584712761019"></a><a name="b584712761019"></a>archive</strong>, see <a href="archiving-log-files.md">Archiving Log Files</a>.</li><li>For <strong id="b13951425191018"><a name="b13951425191018"></a><a name="b13951425191018"></a>clear</strong>, see <a href="deleting-part-records.md">Deleting Part Records</a>.</li><li>For <strong id="b792744261017"><a name="b792744261017"></a><a name="b792744261017"></a>config</strong>, see <a href="updating-a-configuration-file.md">Updating a Configuration File</a>.</li><li>For <strong id="b1469116919117"><a name="b1469116919117"></a><a name="b1469116919117"></a>help</strong>, see <a href="viewing-command-help-information.md">Viewing Command Help Information</a>.</li><li>For <strong id="b5810132917111"><a name="b5810132917111"></a><a name="b5810132917111"></a>version</strong>, see <a href="querying-the-version-number.md">Querying the Version Number</a>.</li></ul>
</td>
</tr>
</tbody>
</table>

## Running Example<a name="section15899161919244"></a>

-   Take the Windows OS as an example. Run the  **obsutil help mb**  command to view the help information about the command for creating a bucket.

```
obsutil help mb

Summary:
create a bucket with the specified parameters

Syntax:
  obsutil mb obs://bucket [-acl=xxx] [-sc=xxx] [-location=xxx] [-config=xxx]

Options:
  -acl=xxx
    the ACL of the bucket, possible values are [private|public-read|public-read-write]

  -sc=xxx
    the default storage class of the bucket, possible values are: [standard|warm|cold]

  -location=xxx
    the region where the bucket is located

  -config=xxx
    the path to the custom config file when running this command
```

