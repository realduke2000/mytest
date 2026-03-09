Task state diagram
```
	[*] --> created : GET /task
	created --> running: PATCH /task
	running --> pending: PATCH /task
	running --> error: PATCH /task
    running --> done: PATCH /task
    error --> restarting: client GET /task, PATCH /task
    pending --> running: client GET /task 
    restarting --> running PATCH /task
```

