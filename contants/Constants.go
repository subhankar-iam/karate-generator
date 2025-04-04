package contants

const (
	User_Prompt = `Context: You are an intelligent Software Tester having knowledge of BDD, karate Framework, Kafka ,MongoDb and other software technologies
                   given the following data in json format %s , where 	background fields says about the background we need to set for the karate
                   feature , the config field in background says if any config file needed to be imported like if mongodb connection string or kafka config needed to be read from some location,
				   then the additional task needed to be performed before scenarios for all the config fetching and additional task (if any) write inline function according to 
				   karate standards. the annotation filed contains value, please put the annotations on top of each scenarios not  background. The header can be a headerFile location, or header in json string.
                   then the Scenario array each element defines one bdd scenarios, in which the request defines the request body in string or the path of the request file should be in format 'classpath:<request filepath>'. additional_instructions refers
                   if any other tasks needed to be done, for those instructions please write karate inline function. other fields are self explanatory.

                   your task is to read the above mentioned json data with the way explained analyse it and generate one karate feature file in gherkin language following
                   the  best practices. Please try to understand the requirements first and then write the executable tests with the data given dont assume anything of your own. 
				   you might need to write some inline functions based on some instructions if it is so please understand the user requirements and write the feature files.
 
                   DISCLAIMER: Alert !! Don't do anything more than instructed and Stricly only send the output in gherkin format and follow the output gherkin format and dont assume anything on your own, No explanation , warning
                   nothing extra noise needed only send the updated gherkin format 

					`

	Gemini_url     = `https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent?key=%s`
	Gemini_Request = `{
                       "contents": {
                         "role": "user",
                         "parts": [
                           {
                             "text": "%s"
                           }
                         ]
                        }
                  }`
)
