# ResponseCreatePazeDigitalWalletSession

Successful Response


## Supported Types

### PazeWebSession

```go
responseCreatePazeDigitalWalletSession := operations.CreateResponseCreatePazeDigitalWalletSessionPazeWebSession(components.PazeWebSession{/* values here */})
```

### PazeMobileSession

```go
responseCreatePazeDigitalWalletSession := operations.CreateResponseCreatePazeDigitalWalletSessionPazeMobileSession(components.PazeMobileSession{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch responseCreatePazeDigitalWalletSession.Type {
	case operations.ResponseCreatePazeDigitalWalletSessionTypePazeWebSession:
		// responseCreatePazeDigitalWalletSession.PazeWebSession is populated
	case operations.ResponseCreatePazeDigitalWalletSessionTypePazeMobileSession:
		// responseCreatePazeDigitalWalletSession.PazeMobileSession is populated
}
```
