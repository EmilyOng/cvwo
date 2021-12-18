import { useState, useEffect } from 'react'
import { BoardPrimitive } from 'types/board'
import { BoardAPI } from 'api/board'

export function useSelectableBoards() {
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')

  const api = new BoardAPI()
  const [boards, setBoards] = useState<BoardPrimitive[]>([])

  function updateBoards(boards: BoardPrimitive[]) {
    setBoards(boards)
  }

  useEffect(() => {
    setLoading(true)
    api
      .getBoards()
      .then((res) => {
        if (res.error) {
          setError(res.error)
        } else {
          setBoards(res)
        }
      })
      .finally(() => setLoading(false))
    return () => {
      // Clean-up
      setLoading(false)
      setError('')
      setBoards([])
    }
  }, [])

  return {
    loading,
    error,
    boards,
    updateBoards
  }
}
