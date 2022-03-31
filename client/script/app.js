const URL = `http://localhost:10020`

/**
 * Fonction d'appel http avec le serveur
 * 
 * @param {String} url 
 * @param {Object} params 
 * @return {Promise<Object>}
 */
const json_fetch = async (url, params = {}) => {
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
      if (error) {
        return Promise.reject(error)
      }

      return data
   }
   return Promise.reject(new Error(`no response for query ${url}`))
}

/**
 * transforme l'heure sous forme de string en la valeur numérique pour 
 * le placement dans le tableau
 * 
 * @param {string} t heure sur 24 heures ex: 13:30
 * @returns {number}
 */
const time_to_int = (t) => ((+t.substring(0, 2) - 2) + (t.substring(3, 5) === '30' ? 2 : 0))

/**
 * gestions des évènements
 */
const events = {
  container: document.querySelector('#calendar--body'),
  /**
   * 
   * @param {Object} { at, from, to } informations relatives è l'évènement
   */
  make_event: function({ at, from, to }) {
    // génération du text
    const text = document.createElement('span')
    text.textContent = `${from} - ${to}`

    // génération du container
    const container = document.createElement('div')
    container.appendChild(text)

    const node = document.createElement('div')
    node.classList.add('event')

    from = time_to_int(from)
    to = time_to_int(to)

    node.style = `--col: ${(new Date(at).getDay()) * 2}; --row-start: ${from}; --duration: ${to-from}`
    node.appendChild(container)

    return node
  },
  /**
   * retire les événements présent dans le calendrier et génères 
   * les événements de la semaine représentée
   */
  generate: function() {
    this.remove()

    // TODO from - to get information sur semaine représentée par calendrier
    json_fetch(`${URL}/events`, { method: 'POST', body: { from: null, to: null } })
      .then((data) => {
        console.log(data)
        data.events && data.events.forEach(
          (event) => this.container.appendChild(this.make_event(event))
        )
      })
      .catch((e) => console.log(e))
  },
  /**
   * retire les événements présents dans le calendrier
   */
  remove: function() {
    this.container.querySelectorAll('.event').forEach((el) => this.container.removeChild(el))
  }
}



events.generate()