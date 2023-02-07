package cmd

import (
	_ "errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"golang.org/x/sys/windows/registry"

	"time"
	"github.com/spf13/cobra"
)



var (
	validArgs  = []string{"key"}
	argAliases = []string{"key"}
)

var productKey string;


func amAdmin() bool {
    _, err := os.Open("\\\\.\\PHYSICALDRIVE0")
    if err != nil {
        return false
    }
    return true
}

func getProductName()  string{
	key, _ := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE) // error discarded for brevity
	defer key.Close()
	productName, _, _ := key.GetStringValue("ProductName") // error discarded for brevity
	return productName;
}

func activateWindows11(winKey string) error {

	registerWinKeyCmd := "slmgr.vbs /ipk " + winKey

	//connect to my KMS server
	conectKMSServercmd := "slmgr.vbs /skms kms8.msguides.com"

	activateDevice := "slmgr.vbs /ato"


	fmt.Println("Windows 11 Activator -): Registering product key for activation...")
	fmt.Println("")
	cmd := exec.Command("cmd", "/C", registerWinKeyCmd)
	// The `Output` method executes the command and
	// collects the output, returning its value
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	errc := cmd.Run()
	if errc != nil {
		return fmt.Errorf("Windows 11 Activator -): Unable to register product key for activation: %v", errc)
	}


	time.Sleep(time.Second)
	fmt.Println("Windows 11 Activator -): Registering Windows 11 to KMS Server...")
	fmt.Println("")
	cmd2 := exec.Command("cmd", "/C", conectKMSServercmd)
	cmd2.Stdout = os.Stdout
	cmd2.Stderr = os.Stderr
	errcc := cmd2.Run()
	if errc != nil {
		return fmt.Errorf("Windows 11 Activator -): Unable to connect to any KMS Server: %v", errcc)
	}


	time.Sleep(time.Second)
	fmt.Println("Windows 11 Activator -): Activating windows 11...")
	fmt.Println("")
	cmd3 := exec.Command("cmd", "/C", activateDevice)
	cmd3.Stdout = os.Stdout
	cmd3.Stderr = os.Stderr
	errccc := cmd3.Run()
	if errc != nil {
		return fmt.Errorf("Windows 11 Activator -): Unable to activate windows 11: %v", errccc)
	}

	return nil
}

//command for scaffolding nextjs application
var activateCmd = &cobra.Command{
	Use:     "activate",
	Short:   "Windows 11 Activator for Pro, Workstation, Home, Education, Enterprise, Core editions",
	Aliases: []string{"s"},
	//validation for argument
	// Args: func(cmd *cobra.Command, args []string) error {
	// 	if len(args) < 1 {
	// 		return errors.New("requires product key to activate windows 11 os")
	// 	}
	// 	return nil
	// },
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Checking system requirements... \n\n")
		if runtime.GOOS != "windows" {
			panic("windows 11 Activator Error -):  can only run on windows for now")
		}

		//check if user is admin
		if !amAdmin() {
			panic("windows 11 Activator Error -): You need to run Windows 11 Activator as an administrator")
		}

	},
	Run: func(cmd *cobra.Command, args []string) {


		fmt.Printf("Kindly take a cofee while your  "+ getProductName() +" is activated...\n")
		fmt.Println("")

		errorx := activateWindows11(productKey)
		fmt.Println(errorx)
		if errorx != nil {
			return
		} else {
			fmt.Printf("Congratulations, your " + getProductName() + " has been activated!\n")
		}

	},
	ValidArgs:  validArgs,
	ArgAliases: argAliases,
}

func init() {

	rootCmd.AddCommand(activateCmd)
	activateCmd.Flags().StringVarP(&productKey, "key", "n", "", "Windows 11 Product Key")

	//Suggestions when “unknown command” happens
	rootCmd.DisableSuggestions = false
	rootCmd.SuggestionsMinimumDistance = 1

	//ensure that product key is a reqired flag
	if err := activateCmd.MarkFlagRequired("key"); err != nil {
		fmt.Println(err)
	}

}