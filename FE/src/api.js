const BASE_URL = import.meta.env.VITE_API_URL || ''

export async function getMessage(signal) {
  const res = await fetch(`${BASE_URL}/api/message`, { signal })
  if (!res.ok) throw new Error(`Request failed with status ${res.status}`)
  const data = await res.json()
  return data.message
}
