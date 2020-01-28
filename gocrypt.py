"""
An example of building a Python library on top of an existing Go package, using Cgo
"""
from ctypes import cdll, c_char_p


class Encryption:
    """
    An example of building a Python library on top of an existing Go package, using Cgo
    """

    def __init__(self):
        """ Constructor for the Encryption class """
        self.crypto = cdll.LoadLibrary("lib/libgocrypt.so")
        self.crypto.Encrypt.argtypes = [c_char_p, c_char_p]
        self.crypto.Encrypt.restype = c_char_p
        self.crypto.Decrypt.argtypes = [c_char_p, c_char_p]
        self.crypto.Decrypt.restype = c_char_p

    def encrypt_string(self, payload, passphrase: str) -> str:
        """ Given a plaintext payload and a passphrase, encrypts the payload """
        encrypted_value = self.crypto.Encrypt(payload.encode(), passphrase.encode())
        return str(encrypted_value, "utf-8")

    def decrypt_string(self, payload, passphrase: str) -> str:
        """ Given an obfuscated payload and passphrase, decrypts the payload """
        decrypted_value = self.crypto.Decrypt(payload.encode(), passphrase.encode())
        return str(decrypted_value, "utf-8")
