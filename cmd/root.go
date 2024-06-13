/*
Copyright © 2024 Kaleb Hawkins <KalebHawkins@outlook.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/spf13/cobra"
)

var (
	region  string
	keyId   string
	outfile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "awskp",
	Short: "Retrieve AWS Public/Private Key from an EC2 keypair.",
	Long: `awskp is used retrieve the public and private key
	from a keypair created using AWS CDK.
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String(region),
		})
		if err != nil {
			return err
		}

		privkey, err := getPrivateKey(sess)
		if err != nil {
			return err
		}

		if outfile != "" {
			err = os.WriteFile(outfile, []byte(privkey), 0600)
			if err != nil {
				return err
			}
		} else {
			fmt.Println(privkey)
		}

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.awskp.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.Flags().StringVarP(&region, "region", "r", "", "The region the keypair resides in")
	rootCmd.MarkFlagRequired("region")
	rootCmd.Flags().StringVarP(&keyId, "key-id", "k", "", "The name of the key pair")
	rootCmd.MarkFlagRequired("key-id")
	rootCmd.Flags().StringVarP(&outfile, "outfile", "o", "", "The file to output the private key")
}

func getPrivateKey(session *session.Session) (string, error) {
	svc := ssm.New(session)

	params := &ssm.GetParameterInput{
		Name:           aws.String(fmt.Sprintf("/ec2/keypair/%s", keyId)),
		WithDecryption: aws.Bool(true),
	}

	resp, err := svc.GetParameter(params)
	if err != nil {
		return "", err
	}

	privKey := *resp.Parameter.Value
	return privKey, nil
}
