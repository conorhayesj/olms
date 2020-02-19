import React, { Component } from "react";
import Books from "./components/books"

let endpoint = "http://localhost:8090";

class Olms extends Component {

	constructor(props){
		super(props);
		this.state = {
			books: [],
			name: ""
		};
	}

	componentDidMount() {
		fetch(endpoint + "/api/v1/allBooks")
			.then(res => res.json())
			.then((data) => {
				this.setState({ books: data })
			})
			.catch(console.log)
	}

	handleChange = (e) => {
		this.setState({[e.target.name]: e.target.value});
	}
	handleSubmit = (e) => {
	//	e.preventDefault();
		const data = { name : this.state.name }
		console.log(data)
		fetch(endpoint + "/api/v1/addBook", {
			mode: "no-cors",
			credentials: "same-origin",
			method: "POST",
			headers: {
				//"Content-Type": "application/x-www-form-urlencoded"
				"Content-Type": "application/json"
			},
			body: JSON.stringify(data)
		}).then(function(response) {
			console.log(response)
			return response;
		});
	}

	render() {
		return (
			<div>
				<div className="row">
					<form>
						<input onChange={this.handleChange} name="name" type="text" id="name" value={this.state.name} placeholder="Enter new book name"/>
						<button onClick={this.handleSubmit}>Submit</button>
					</form>
				</div>
					
			<Books books={this.state.books} />
			</div>
		)
	}
}

export default Olms;
