import NextAuth from 'next-auth';
import CredentialsProvider from 'next-auth/providers/credentials';

const handler = NextAuth({
  providers: [
    CredentialsProvider({
      name: 'Credentials',
      credentials: {
        email: { label: "Email", type: "email" },
        password: { label: "Password", type: "password" }
      },
      async authorize(credentials) {
        // This is where you would typically verify the user credentials
        // against your database
        if (credentials?.email === "user@example.com" && credentials?.password === "password") {
          return {
            id: "1",
            email: credentials.email,
            name: "Example User",
          }
        }
        return null;
      }
    })
  ],
  session: {
    strategy: "jwt",
  },
  pages: {
    signIn: '/login',
  },
  theme: {
    colorScheme: "dark",
  },

});

export { handler as GET, handler as POST };
