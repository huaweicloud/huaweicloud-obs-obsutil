# Using the obsutil help Command to Search for Functions<a name="EN-US_TOPIC_0190483383"></a>

obsutil provides  **help**  commands for viewing the help documents of each command. To query the help document of the bucket creation command, perform the following steps:

1.  Run the  **obsutil help**  command to query the list of all supported commands.
2.  Find the abbreviation of the command to be viewed based on the document description in the command list. For example, the abbreviation of the command for creating a bucket is  **mb**.
3.  Run the  **obsutil help mb**  command to view the usage and detailed functions of the  **mb**  command, illustrated as follows:

    ```
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

4.  Run the  **obsutil mb obs://bucket-test -location xxx**  command to create a bucket named  **bucket-test**  in the  _xxx_  region.

>![](public_sys-resources/icon-note.gif) **NOTE:**   
>-   For more information about the  **help**  command, see  [Viewing Command Help Information](viewing-command-help-information.md).  
>-   You can set the  **helpLanguage**  parameter in the configuration file to configure the language type of the  **help**  command. For example,  **helpLanguage=Chinese**  indicates that the language type of the help command is Chinese.  
>-   The supported languages are Chinese and English. The default language is English.  

