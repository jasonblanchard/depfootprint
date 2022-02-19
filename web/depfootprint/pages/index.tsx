import { useRef } from 'react';
import type { NextPage } from 'next'
import { useRouter } from 'next/router';

const Home: NextPage = () => {
  const inputEl = useRef<HTMLInputElement>(null);
  const router = useRouter();

  function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();
    const value = inputEl?.current?.value;
    if (!value) return;
    router.push(`/packages/${value}`);
  }

  return (
    <div className="container mx-auto my-5">
      <form onSubmit={handleSubmit}>
        <label>
          NPM package:
          <input placeholder="express" ref={inputEl} className="border-2 rounded border-neutral-600 mx-2 p-1" type="text" />
        </label>
        <button className="border-2 border-neutral-600 rounded px-4 py-1 bg-neutral-300 hover:bg-neutral-400 focus:bg-neutral-400">Go</button>
      </form>
    </div>
    )
}

export default Home
