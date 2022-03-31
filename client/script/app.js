/**
 * Fonction d'appel http avec le serveur
 * 
 * @param {String} url 
 * @param {Object} params 
 * @return {Promise<Object>}
 */
const jsonFetch = async (url, params = {}) => {
  if (params.body && typeof params.body === 'object') {
    params.body = JSON.stringify(params.body)
  }

  params = {
    headers: {
      'Content-Type': 'application/json',
      'Accept': 'application/json',
    },
    ...params
  }

  const response = await fetch(url, params)
   if (response.ok) {
      const { data, error } = await response.json()
      if (error !== null) {
        return Promise.reject(error)
      }

      return data
   }
   return Promise.reject(new Error(`no response for query ${url}`))
}
