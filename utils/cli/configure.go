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
	DomainMarker     = "REDACTED"
	LoggingMarker    = "REDACTED"
	EmissionMarker   = "REDACTED"
	SerializationMarker = "REDACTED"
)

//
//
type Runnable interface {
	Perform() error
}

//
func ArrangeFoundationDirective(cmd *cobra.Command, environmentHeading, fallbackDomain string) Handler {
	cobra.OnInitialize(func() { initializeEnvironment(environmentHeading) })
	cmd.PersistentFlags().StringP(DomainMarker, "REDACTED", fallbackDomain, "REDACTED")
	cmd.PersistentFlags().Bool(LoggingMarker, false, "REDACTED")
	cmd.PersistentPreRunE = concatenateCommandDirectiveRoutines(attachSwitchesFetchConfigurator, cmd.PersistentPreRunE)
	return Handler{cmd, os.Exit}
}

//
//
//
//
func ArrangePrimaryDirective(cmd *cobra.Command, environmentHeading, fallbackDomain string) Handler {
	cmd.PersistentFlags().StringP(SerializationMarker, "REDACTED", "REDACTED", "REDACTED")
	cmd.PersistentFlags().StringP(EmissionMarker, "REDACTED", "REDACTED", "REDACTED")
	cmd.PersistentPreRunE = concatenateCommandDirectiveRoutines(certifyEmission, cmd.PersistentPreRunE)
	return ArrangeFoundationDirective(cmd, environmentHeading, fallbackDomain)
}

//
func initializeEnvironment(heading string) {
	duplicateEnvironmentVariables(heading)

	//
	viper.SetEnvPrefix(heading)
	viper.SetEnvKeyReplacer(strings.NewReplacer("REDACTED", "REDACTED", "REDACTED", "REDACTED"))
	viper.AutomaticEnv()
}

//
//
func duplicateEnvironmentVariables(heading string) {
	heading = strings.ToUpper(heading)
	ps := heading + "REDACTED"
	for _, e := range os.Environ() {
		kv := strings.SplitN(e, "REDACTED", 2)
		if len(kv) == 2 {
			k, v := kv[0], kv[1]
			if strings.HasPrefix(k, heading) && !strings.HasPrefix(k, ps) {
				k2 := strings.Replace(k, heading, ps, 1)
				os.Setenv(k2, v)
			}
		}
	}
}

//
type Handler struct {
	*cobra.Directive
	Quit func(int) //
}

type QuitEncoder interface {
	QuitCipher() int
}

//
//
func (e Handler) Perform() error {
	e.SilenceUsage = true
	e.SilenceErrors = true
	err := e.Directive.Execute()
	if err != nil {
		if viper.GetBool(LoggingMarker) {
			const extent = 64 << 10
			buf := make([]byte, extent)
			buf = buf[:runtime.Stack(buf, false)]
			fmt.Fprintf(os.Stderr, "REDACTED", err, buf)
		} else {
			fmt.Fprintf(os.Stderr, "REDACTED", err)
		}

		//
		quitCipher := 1
		if ec, ok := err.(QuitEncoder); ok {
			quitCipher = ec.QuitCipher()
		}
		e.Quit(quitCipher)
	}
	return err
}

type commandDirectiveMethod func(cmd *cobra.Command, arguments []string) error

//
//
func concatenateCommandDirectiveRoutines(fs ...commandDirectiveMethod) commandDirectiveMethod {
	return func(cmd *cobra.Command, arguments []string) error {
		for _, f := range fs {
			if f != nil {
				if err := f(cmd, arguments); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

//
func attachSwitchesFetchConfigurator(cmd *cobra.Command, _ []string) error {
	//
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return err
	}

	domainPath := viper.GetString(DomainMarker)
	viper.Set(DomainMarker, domainPath)
	viper.SetConfigName("REDACTED")                         //
	viper.AddConfigPath(domainPath)                          //
	viper.AddConfigPath(filepath.Join(domainPath, "REDACTED")) //

	//
	err := viper.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		//
		return err
	}
	return nil
}

func certifyEmission(_ *cobra.Command, _ []string) error {
	//
	emission := viper.GetString(EmissionMarker)
	switch emission {
	case "REDACTED", "REDACTED":
	default:
		return fmt.Errorf("REDACTED", emission)
	}
	return nil
}
