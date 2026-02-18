package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	HomeMark     = "REDACTED"
	TrackMark    = "REDACTED"
	ResultMark   = "REDACTED"
	CodecMark = "REDACTED"
)

//
//
type Runnable interface {
	Perform() error
}

//
func ArrangeRootCommand(cmd *cobra.Command, contextPrefix, standardHome string) Runner {
	cobra.OnInitialize(func() { initContext(contextPrefix) })
	cmd.PersistentFlags().StringP(HomeMark, "REDACTED", standardHome, "REDACTED")
	cmd.PersistentFlags().Bool(TrackMark, false, "REDACTED")
	cmd.PersistentPreRunE = joinCobraCommandRoutines(attachMarksImportViper, cmd.PersistentPreRunE)
	return Runner{cmd, os.Exit}
}

//
//
//
//
func ArrangeMainCommand(cmd *cobra.Command, contextPrefix, standardHome string) Runner {
	cmd.PersistentFlags().StringP(CodecMark, "REDACTED", "REDACTED", "REDACTED")
	cmd.PersistentFlags().StringP(ResultMark, "REDACTED", "REDACTED", "REDACTED")
	cmd.PersistentPreRunE = joinCobraCommandRoutines(certifyResult, cmd.PersistentPreRunE)
	return ArrangeRootCommand(cmd, contextPrefix, standardHome)
}

//
func initContext(prefix string) {
	cloneContextVars(prefix)

	//
	viper.SetEnvPrefix(prefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer("REDACTED", "REDACTED", "REDACTED", "REDACTED"))
	viper.AutomaticEnv()
}

//
//
func cloneContextVars(prefix string) {
	prefix = strings.ToUpper(prefix)
	ps := prefix + "REDACTED"
	for _, e := range os.Environ() {
		kv := strings.SplitN(e, "REDACTED", 2)
		if len(kv) == 2 {
			k, v := kv[0], kv[1]
			if strings.HasPrefix(k, prefix) && !strings.HasPrefix(k, ps) {
				k2 := strings.Replace(k, prefix, ps, 1)
				os.Setenv(k2, v)
			}
		}
	}
}

//
type Runner struct {
	*cobra.Directive
	Quit func(int) //
}

type QuitEncoder interface {
	QuitCode() int
}

//
//
func (e Runner) Perform() error {
	e.SilenceUsage = true
	e.SilenceErrors = true
	err := e.Directive.Execute()
	if err != nil {
		if viper.GetBool(TrackMark) {
			const volume = 64 << 10
			buf := make([]byte, volume)
			buf = buf[:runtime.Stack(buf, false)]
			fmt.Fprintf(os.Stderr, "REDACTED", err, buf)
		} else {
			fmt.Fprintf(os.Stderr, "REDACTED", err)
		}

		//
		quitCode := 1
		if ec, ok := err.(QuitEncoder); ok {
			quitCode = ec.QuitCode()
		}
		e.Quit(quitCode)
	}
	return err
}

type cobraCommandFunction func(cmd *cobra.Command, args []string) error

//
//
func joinCobraCommandRoutines(fs ...cobraCommandFunction) cobraCommandFunction {
	return func(cmd *cobra.Command, args []string) error {
		for _, f := range fs {
			if f != nil {
				if err := f(cmd, args); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

//
func attachMarksImportViper(cmd *cobra.Command, _ []string) error {
	//
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return err
	}

	homeFolder := viper.GetString(HomeMark)
	viper.Set(HomeMark, homeFolder)
	viper.SetConfigName("REDACTED")                         //
	viper.AddConfigPath(homeFolder)                          //
	viper.AddConfigPath(filepath.Join(homeFolder, "REDACTED")) //

	//
	err := viper.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		//
		return err
	}
	return nil
}

func certifyResult(_ *cobra.Command, _ []string) error {
	//
	result := viper.GetString(ResultMark)
	switch result {
	case "REDACTED", "REDACTED":
	default:
		return fmt.Errorf("REDACTED", result)
	}
	return nil
}
