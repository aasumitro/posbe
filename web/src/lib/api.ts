const API_URL = `http://localhost:8000/api/v1`

const Endpoint = {
  Auth: {},
  User: { },
  Catalog: {
    Addon: "addons"
  }
}

const get = async (path: string)  => {
  try {
    const url = `${API_URL}/${path}`
    const response = await fetch(url, {
      method: "GET",
      headers: {'Content-Type': 'application/json'},
      credentials: 'include',
    })
    const content = await response.json();
    return Promise.resolve(content)
  } catch (e) {
    return Promise.reject(e)
  }
}


export {
  Endpoint,
  get,
}
