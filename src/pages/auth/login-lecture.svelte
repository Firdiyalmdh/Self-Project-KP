<script>
  import { redirect, url } from "@roxi/routify";
  import { getData, postFormData } from "../../lib/request";
  import { user } from "../../stores/UserStore";
      
  let email ='', pass='';

$:login = async () => {
    postFormData({
        endpoint: '/login',
        payload: {
            email,
            pass
        },
        onSuccess: (response) => {
            getSession(response.InsertedID)
        },
        onFailed: (response) => {
            console.log(response);
            alert(response.data.message)
        }
    })
}

$:getSession = async (id) => {
    getData({
      endpoint: `/session/${id}`,
      onSuccess: (response) => { 
        user.set({
          id: response.json,
          sessionId: response._id,
          name: response.nama,
          role: response.role
        })
        $redirect('/lecture/home')
      },
      onFailed: (response) => {
          console.log(response);
          alert(response.data.message)
      }
    })
  }

</script>

<div class="font-poppins bg-login w-full flex justify-center align-center">
<form on:submit|preventDefault={login} class="bg-white h-96 w-96 shadow-md rounded-3xl px-8 pt-6 pb-8 mb-4">
  <div class="grid grid-cols-2 mb-5" style="max-width: 25%;">
    <img src="/images/pens-color.png">
    <img src="/images/hmce-color.png">
  </div>
  <div class="text-sm font-medium">Welcome to</div>
  <p class="text-4xl font-extrabold">CEMIS - Lecture</p>
  <div class="text-[0.6rem] mt-5 mb-4">
    <input bind:value={email}
      class="h-10 font-thin border rounded-lg w-full py-2 px-3 text-gray-400 leading-tight focus:outline-1 outline-blue-300 focus:shadow-outline" 
      id="username" 
      type="text" 
      placeholder=" Enter your email">
  </div>
  <div class="text-[0.6rem] mt-5 mb-4">
    <input bind:value={pass} 
      class="h-10 border rounded-lg w-full py-2 px-3 text-gray-400 leading-tight focus:outline-1 outline-blue-300 focus:shadow-outline" 
      id="password" 
      type="password" 
      placeholder="Password">
    <p class="mt-1 text-end text-blue-500">Forgot Password</p>
  </div>
  <button class="text-[0.7rem] text-white bg-blue-500 w-full rounded-lg h-10 shadow-sm shadow-gray-300">Login</button>
</form>
</div>