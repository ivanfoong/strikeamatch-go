package strikeamatch_test

import (
  "testing"
  "github.com/ivanfoong/strikeamatch-go/strikeamatch"
  "fmt"
)

func TestCompareString(t *testing.T) {
  var testCases = []struct {
    inputs []string
    output float64
  }{
    {[]string{"Healed", "Sealed"}, 0.8},
    {[]string{"Healed", "Healthy"}, 0.55},
    {[]string{"Healed", "Heard"}, 0.44},
    {[]string{"Healed", "Herded"}, 0.4},
    {[]string{"Healed", "Help"}, 0.25},
    {[]string{"Healed", "Sold"}, 0.0},
    {[]string{"Web Database Applications with PHP & MySQL", "Web Database Applications"}, 0.82},
    {[]string{"Web Database Applications with PHP & MySQL", "PHP Web Applications"}, 0.68},
    {[]string{"Web Database Applications with PHP & MySQL", "Web Aplications"}, 0.59},
    {[]string{"Creating Database Web Applications with PHP and ASP", "Web Database Applications"}, 0.71},
    {[]string{"Creating Database Web Applications with PHP and ASP", "PHP Web Applications"}, 0.59},
    {[]string{"Creating Database Web Applications with PHP and ASP", "Web Aplications"}, 0.5},
    {[]string{"Building Database Applications on the Web Using PHP3", "Web Database Applications"}, 0.7},
    {[]string{"Building Database Applications on the Web Using PHP3", "PHP Web Applications"}, 0.58},
    {[]string{"Building Database Applications on the Web Using PHP3", "Web Aplications"}, 0.49},
    {[]string{"Building Web Database Applications with Visual Studio 6", "Web Database Applications"}, 0.67},
    {[]string{"Building Web Database Applications with Visual Studio 6", "PHP Web Applications"}, 0.47},
    {[]string{"Building Web Database Applications with Visual Studio 6", "Web Aplications"}, 0.46},
    {[]string{"Web Application Development With PHP", "Web Database Applications"}, 0.51},
    {[]string{"Web Application Development With PHP", "PHP Web Applications"}, 0.67},
    {[]string{"Web Application Development With PHP", "Web Aplications"}, 0.56},
    {[]string{"WebRAD: Building Database Applications on the Web with Visual FoxPro and Web Connection", "Web Database Applications"}, 0.49},
    {[]string{"WebRAD: Building Database Applications on the Web with Visual FoxPro and Web Connection", "PHP Web Applications"}, 0.34},
    {[]string{"WebRAD: Building Database Applications on the Web with Visual FoxPro and Web Connection", "Web Aplications"}, 0.33},
    {[]string{"Structural Assessment: The Role of Large and Full-Scale Testing", "Web Database Applications"}, 0.13},
    {[]string{"Structural Assessment: The Role of Large and Full-Scale Testing", "PHP Web Applications"}, 0.07},
    {[]string{"Structural Assessment: The Role of Large and Full-Scale Testing", "Web Aplications"}, 0.07},
    {[]string{"How to Find a Scholarship Online", "Web Database Applications"}, 0.1},
    {[]string{"How to Find a Scholarship Online", "PHP Web Applications"}, 0.11},
    {[]string{"How to Find a Scholarship Online", "Web Aplications"}, 0.12},
  }
  for _, tt := range testCases {
    text1 := tt.inputs[0]
    text2 := tt.inputs[1]
    expectedScore := tt.output
    score := strikeamatch.CompareString(text1, text2)
    if fmt.Sprintf("%.2f", expectedScore) != fmt.Sprintf("%.2f", score) {
      t.Errorf("expected CompareString(%s, %s) to equal `%.2f`, got `%.2f` instead", text1, text2, expectedScore, score)
    }
  }
}