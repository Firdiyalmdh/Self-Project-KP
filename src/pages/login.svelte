<script>
    import App from '../App.svelte';
</script>

<style></style>

<div class ="bg-sky-600 h-screen flex justify-center align-center">  
  <div class="border-2 border-black bg-amber-200 w-96 self-center flex-initialshadow-md rounded px-8 pt-6 pb-8 mb-4 flex flex-col">
    <img class="-mt-20" src="./images/BannerSignIn.png">  
    <div class="-mt-5 mb-2 ml-5 mr-5">
        <div class="mb-2 h-7 flex border-2 border-black rounded-3xl bg-amber-400 shadow appearance-none w-full px-3">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-4 ">
                <path stroke-linecap="round" d="M16.5 12a4.5 4.5 0 11-9 0 4.5 4.5 0 019 0zm0 0c0 1.657 1.007 3 2.25 3S21 13.657 21 12a9 9 0 10-2.636 6.364M16.5 12V8.25" />
              </svg>              
            <input class=" text-grey-darker bg-transparent" id="email" type= "text" placeholder="Email">
        </div>
    </div>
    <div class="mb-2 ml-5 mr-5">
      <div class="mb-2 h-7 flex border-2 border-black rounded-3xl bg-amber-400 shadow appearance-none w-full px-3">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-4">
          <path stroke-linecap="round" stroke-linejoin="round" d="M16.5 10.5V6.75a4.5 4.5 0 10-9 0v3.75m-.75 11.25h10.5a2.25 2.25 0 002.25-2.25v-6.75a2.25 2.25 0 00-2.25-2.25H6.75a2.25 2.25 0 00-2.25 2.25v6.75a2.25 2.25 0 002.25 2.25z" />
        </svg>               
          <input class=" text-grey-darker bg-transparent" id="passsword" type= "text" placeholder="Password">
      </div>
  </div>
  
  <button class="bg-sky-600 hover:bg-sky-800 text-black font-bold py-1 px-4 rounded-3xl border-2 border-black ml-20 mr-20 ">Login</button>
  <div class="flex items-center justify-between">
  </div>
  <p class="mt-2 flex justify-center font-bold align-center text-xs">New to Whatido? Sign up for free</p>
    <button class="mt-3 bg-transparent">
        <img class="inline-flex items-center w-20" src="./images/ButtonSignUp.png">
    </button>
  </div>
</div>