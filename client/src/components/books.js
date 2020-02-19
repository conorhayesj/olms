import React from "react"

const Books = ({books}) => {
	return (
		<div>
			<center><h1>All Books</h1></center>
			{books.map((book) => (
				<div class="card">
					<div class="card-body">
						<h5 class="card-title">{book.name}</h5>
					</div>
				</div>
			))}
		</div>
	)
};

export default Books
