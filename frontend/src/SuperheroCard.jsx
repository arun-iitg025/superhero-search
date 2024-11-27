import React from "react";

const SuperheroCard = ({ superhero }) => {
  return (
    <div className="movie">
    <div>
        <p>{superhero.powers}</p>
      </div>
      <div>
        <img src={superhero.image} alt={superhero.name} />
      </div>
      <div>
        <h3>{superhero.name}</h3>
      </div>


      {/* <div>
        {/* Display the superhero's image *
        <img
          src={superhero.image || "https://via.placeholder.com/150"}
          alt={superhero.name}
        />
      </div>
      <div>
        <h3>{superhero.name}</h3>
        <p><strong>Powers:</strong> {superhero.powers?.join(", ")}</p>
        <p><strong>Movies:</strong> {superhero.movies?.join(", ")}</p>
      </div> */}
    </div>
  );
};

export default SuperheroCard;
