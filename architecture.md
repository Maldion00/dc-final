Architecture Document
=====================
Architecture                                                             

					                    ______> Worker
	 JSON	      MSG		   Jobs     |
Client --------> API ------> Controller -----------> Scheduler |-----> Worker
							            |_____> Worker
							              RPC