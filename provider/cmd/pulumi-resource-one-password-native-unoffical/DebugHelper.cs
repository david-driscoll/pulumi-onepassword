using System.Diagnostics;

namespace pulumi_resource_one_password_native_unoffical;

public static class DebugHelper
{
    public static void WaitForDebugger()
    {
        if (Environment.GetCommandLineArgs().Any(z => z == "--debugger" || z == "--debug"))
            while (!Debugger.IsAttached)
            {
                Thread.Sleep(1000);
            }
    }
}