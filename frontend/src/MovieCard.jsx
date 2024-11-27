import React from "react";

const MovieCard = ({ item }) => {
  return (
    <div className="movie">
      <div>
        <p>{item.release_year}</p>
      </div>
      <div>
        <img src={item.poster} alt={item.title} />
      </div>
      <div>
        <h3>{item.title}</h3>
      </div>
    </div>
  );
};

export default MovieCard;
