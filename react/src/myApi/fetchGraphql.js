import axios from "axios";

const endpoint = "https://beta.pokeapi.co/graphql/v1beta";

const fetchGraphQL = async (query, variables, operationName) => {
  const graphqlQuery = {
    operationName: operationName,
    query: query,
    variables: variables,
  };

  try {
    const result = await axios({
      url: endpoint,
      method: "post",
      data: graphqlQuery,
    });
    return await result.data;
  } catch (error) {
    console.log(error);
  }

  
};

export default fetchGraphQL;
