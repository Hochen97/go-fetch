class App extends React.Component {
    render() {
       <Home />
    }
}

class Home extends React.Component {
    render() {
        return (
            <div className="container">
                <div className="col-xs-8 col-xs-offset-2 jumbotron text-center">
                    <h1>Twitter Mosaic Bot</h1>
                    <p>a bot to create mosaics from tweets</p>
                    <p>using golang, react, Auth0, and antogen/gosaic</p>
                    <a onClick={this.authenticate} className="btn btn-primary btn-lg btn-login btn-block">Sign In</a>
                </div>
            </div>
        )
    }
}