import { NowRequest, NowResponse } from '@now/node'
import axios from 'axios'

const handler = async (request: NowRequest, response: NowResponse) => {
  const id = request.query.id
  try {
    const res = await axios.get(`https://api.scratch.mit.edu/projects/${id}`)

    response.status(200).send(res.data)
  } catch (error) {
    response.status(error.response.status).send('Error')
  }
}

export default handler
