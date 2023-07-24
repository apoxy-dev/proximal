const List = async () => {
  return fetch(process.env.REACT_APP_API_HOST + '/v1/logs').then((res) => res.json());
};

export default {
  List,
};
