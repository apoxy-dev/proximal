const List = async () => {
  return fetch(process.env.REACT_APP_API_HOST + '/v1/endpoints').then((res) => res.json());
};

const Create = async (endpoint) => {
  return fetch(process.env.REACT_APP_API_HOST + '/v1/endpoints', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ endpoint }),
  }).then((res) => res.json());
};

const Delete = async (cluster) => {
  return fetch(process.env.REACT_APP_API_HOST + '/v1/endpoints/' + cluster, {
    method: 'DELETE',
  }).then((res) => res.json());
};

const SetDefault = async (endpoint) => {
  endpoint.default_upstream = true;
  return fetch(process.env.REACT_APP_API_HOST + '/v1/endpoints/' + endpoint.cluster, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ endpoint }),
  });
};

export default {
  List,
  Create,
  Delete,
  SetDefault,
};
