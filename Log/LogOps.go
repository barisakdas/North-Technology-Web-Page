package Log

import "log"

func LogJson(page,logType,controllerName, functionName, description, errorDetail string)  {
	errorMessage:= `{
						"Type":` + logType + `,
						"Admin/Site":`+ page +` 
						"Controller_Name":` + controllerName + `
						"Function_Name":` + functionName + `,
						"Description":` + description + `,
						"Error_Detail":` + errorDetail + `
					}`
	log.Fatalln(errorMessage)
}
