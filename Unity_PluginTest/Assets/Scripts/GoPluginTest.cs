using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System.Runtime.InteropServices;

public class GoPluginTest : MonoBehaviour
{

#if UNITY_STANDALONE_WIN
 const string pluginDll = "libgo" ;
#elif UNITY_STANDALONE_OSX
 const string pluginDll = "OSXBundle" ;
#elif UNITY_IOS
 const string pluginDll = "__Internal" ;
#elif UNITY_ANDROID
 const string pluginDll = "go" ;
#endif



	[DllImport (pluginDll)]
	private static extern int go_Add (int a,int b);
    [DllImport (pluginDll)]
    private static extern int go_SendMsg(string p0);
    [DllImport (pluginDll)]
    private static extern void go_SetCallbackFunc(SimpleArgCallback p0);

    private delegate string SimpleArgCallback(string arg);
    private string SimpleArgCallbackUnityFun(string arg) {
        Debug.Log( "received callback with arg:" + arg );
        return "unity echo back:" + arg ;
    }



    // Start is called before the first frame update
    void Start()
    {
        Debug.Log( "test start" ) ;

        SimpleArgCallback sc = new SimpleArgCallback(SimpleArgCallbackUnityFun);
        go_SetCallbackFunc(sc);

        Debug.Log( go_Add(1,2) ) ;

        int ret = go_SendMsg( "send from unity" );
        Debug.Log( "ret of SendMsg " + ret  ) ;
    }

    // Update is called once per frame
    void Update()
    {
        
    }


}
