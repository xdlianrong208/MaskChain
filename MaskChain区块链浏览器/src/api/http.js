import Axios from 'axios'

const axiosIns = Axios.create({
  timeout: 30 * 1000
})
axiosIns.defaults.baseURL = 'http://127.0.0.1:8545'
axiosIns.defaults.contentType = 'application/json'
axiosIns.defaults.validateStatus = function () {
  return true
}
export default {
  axiosIns
}

/** post请求方法
 * @param options
 * @return {Promise}
 */
export function post (options) {
  return new Promise((resolve, reject) => {
    axiosIns(options)
      .then(response => {
        resolve(response)
      })
      .catch(error => {
        reject(error)
      })
  })
};
/** get请求方法
 * @param options
 * @return {Promise}
 */
export function get (options) {
  return new Promise((resolve, reject) => {
    axiosIns(options).then(response => {
      resolve(response)
    })
      .catch(error => {
        reject(error)
      })
  })
}

/** patch请求方法
 * @param options
 * @return {Promise}
 */
export function patch (options) {
  return new Promise((resolve, reject) => {
    axiosIns(options).then(response => {
      resolve(response)
    })
      .catch(error => {
        reject(error)
      })
  })
}
/** put请求方法
 * @param options
 * @return {Promise}
 */
export function put (options) {
  return new Promise((resolve, reject) => {
    axiosIns(options).then(response => {
      resolve(response)
    })
      .catch(error => {
        reject(error)
      })
  })
}
/** delete请求方法
 * @param options
 * @return {Promise}
 */
export function deleted (options) {
  return new Promise((resolve, reject) => {
    axiosIns(options).then(response => {
      resolve(response)
    })
      .catch(error => {
        reject(error)
      })
  })
}
