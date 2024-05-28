import { Component } from "react";

function launcherUI() {
  return (
    <div className="absolute text-center top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2">
      <h1 className="text-8xl font-semibold italic">POSBE</h1>
      <p className="mt-8 text-2xl font-extralight">Point of Sales</p>
      <div className="mt-32 w-4 mx-auto">
        <svg
          aria-hidden="true"
          className="w-8 h-8 mr-2 text-gray-200 animate-spin dark:text-gray-600 fill-blue-600"
          viewBox="0 0 100 101"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z"
            fill="currentColor" />
          <path
            d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z"
            fill="currentFill" />
        </svg>
      </div>
    </div>
  )
}

function authUI() {
  return (
    <section className="bg-gray-50 dark:bg-gray-900">
      <div className="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0">
        <div className="mb-8 text-center">
          <p className="text-2xl font-bold italic text-gray-900 dark:text-white">
            POSBE
          </p>
          <span className="text-medium font-light mt-3 mb-6">
            Point of sales
          </span>
        </div>
        <div
          className="w-full bg-white rounded-lg shadow dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700">
          <div className="p-6 space-y-4 md:space-y-6 sm:p-8">
            <h1 className="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl dark:text-white">
              Sign in to your account
            </h1>
            <form className="space-y-4 md:space-y-6" action="#">
              <div>
                <label htmlFor="username" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Your
                  username</label>
                <input type="text" name="username" id="username"
                  className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-gray-600 focus:border-gray-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-gray-500 dark:focus:border-gray-500"
                  placeholder="@username" required />
              </div>
              <div>
                <label htmlFor="password"
                  className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Password</label>
                <input type="password" name="password" id="password" placeholder="••••••••"
                  className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-gray-600 focus:border-gray-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-gray-500 dark:focus:border-gray-500"
                  required />
              </div>
              <button type="submit"
                className="w-full h-12 text-white bg-gray-700 hover:bg-gray-800 focus:ring-4 focus:outline-none focus:ring-gray-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center"
                onClick={() => {
                  localStorage.setItem("is_logged_in", "true")
                }}>Sign in
              </button>
            </form>
          </div>
        </div>
      </div>
    </section>
  )
}

interface PageState {
  isLoading: boolean
  isLoggedIn: false,
}

function WithLauncher(WrappedComponent: any) {
  return class extends Component {
    constructor(props: any) {
      super(props);
      this.state = {
        isLoading: true,
        isLoggedIn: false,
      } as PageState;
    }
    async componentDidMount() {
      try {
        // await api.userProfile().then(() =. {}).catch(() => {});
        if (localStorage.getItem("is_logged_in")) {
          this.setState({ isLoading: false, isLoggedIn: true })
          return
        }

        setTimeout(() => {
          this.setState({ isLoading: false, isLoggedIn: false })
        }, 500)
      } catch (err) {
        this.setState({ isLoading: false })
      }
    }
    render() {
      const pageState = this.state as PageState;
      // while checking user session, show "loading" message
      if (pageState.isLoading) return launcherUI();
      // while users not authenticate then display the login component
      if (!pageState.isLoggedIn) return authUI();
      // otherwise, show the desired route
      return <WrappedComponent {...this.props} />;
    }
  };
}

export { WithLauncher };
