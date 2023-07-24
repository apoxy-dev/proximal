const List = async () => {
  return fetch(process.env.REACT_APP_API_HOST + '/v1/middlewares').then((res) => res.json());
};

const Create = async (data) => {
  return fetch(process.env.REACT_APP_API_HOST + '/v1/middlewares', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ middleware: data }),
  }).then((res) => res.json());
};

const Delete = async (slug) => {
  return fetch(process.env.REACT_APP_API_HOST + '/v1/middlewares/' + slug, {
    method: 'DELETE',
  }).then((res) => res.json());
};

const TriggerBuild = async (slug) => {
  return fetch(process.env.REACT_APP_API_HOST + '/v1/middlewares/' + slug + '/builds', {
    method: 'POST',
  }).then((res) => res.json());
};

export default {
  List,
  Create,
  Delete,
  TriggerBuild,
};
